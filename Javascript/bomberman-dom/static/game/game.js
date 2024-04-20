import {
  useState,
  init,
  useEventListener,
  node,
} from "../frawnwok/frawnwok.js";
import stats from "./stats.js";
import map from "./map.js";

export const MAP_SIZE = 13;

const limitRange = (num, min, max) => {
  min = min ?? 0;
  max = max ?? MAP_SIZE - 1;
  return Math.min(Math.max(num, min), max);
};

const Game = (ws) => {
  const [keyListener, _] = useState(
    useEventListener({
      target: window,
      event: "keydown",
      callbacks: [handleKeys],
    })
  );
  const [mapData, setMapData] = useState([]);
  const [enemies, setEnemies] = useState([]);
  const [bombs, setBombs] = useState([]);
  const [player, setPlayer] = useState({});

  ws.handleOnUpdate = (data) => {
    setMapData((prev) => {
      prev = data.board;
      return prev;
    });
    if (data.player) {
      setPlayer((prev) => {
        prev.x = data.player.x;
        prev.y = data.player.y;
        prev.bombs = data.player.bombs;
        prev.speed = data.player.speed;
        prev.flame = data.player.flame;
        return prev;
      });
    }
  };

  ws.handleOnStart = (data) => {
    setMapData((prev) => {
      prev = data.board;
      return prev;
    });

    setPlayer((prev) => {
      prev.x = data.player.x;
      prev.y = data.player.y;
      prev.life = data.player.life;
      prev.bombs = data.player.bombs;
      prev.bombsPlaced = data.player.bombsPlaced;
      prev.speed = data.player.speed;
      prev.flame = data.player.flame;
      prev.alive = data.player.alive;
      prev.id = data.player.id + 1;
      prev.name = data.player.name.split("_").pop();
      prev.won = false;
      return prev;
    });

    setEnemies((prev) => {
      data.players.map((enemy) => {
        if (enemy.name.split("_").pop() !== player.name) {
          prev.push({
            name: enemy.name,
            id: enemy.id + 1,
            x: enemy.x,
            y: enemy.y,
          });
        }
      });
      return prev;
    });
  };
  ws.handleOnPlayerUpdate = (data) => {
    if (data.message_text === "died") {
      if (data.sender_uuid === player.name) {
        setPlayer((prev) => {
          prev.x = -1;
          prev.y = -1;
          prev.alive = false;
          prev.life = 0;
          return prev;
        });
      } else {
        setEnemies((prev) => {
          prev.map((enemy) => {
            if (enemy.name.split("_").pop() === data.sender_uuid) {
              enemy.x = -1;
              enemy.y = -1;
              enemy.alive = false;
              enemy.life = 0;
            }
          });
          return prev;
        });
      }
    } else if (data.message_text === "damaged") {
      setPlayer((prev) => {
        prev.life -= 1;
        return prev;
      });
    } else if (data.message_text === "won") {
      setEnemies(() => []);
      setPlayer((prev) => {
        prev.won = true;
        return prev
      })
    } else {
      console.log(data);
    }
  };

  ws.handleOnMove = (data) => {
    console.log("Handling move", data.action);
    if (data.action.name === player.name) {
      setPlayer((prev) => {
        prev.x = data.action.x;
        prev.y = data.action.y;
        return prev;
      });
    } else {
      setEnemies((prev) => {
        prev.map((enemy) => {
          if (enemy.name.split("_").pop() === data.action.name) {
            enemy.x = data.action.x;
            enemy.y = data.action.y;
          }
        });
        return prev;
      });
    }
  };

  ws.handleOnBomb = (data) => {
    console.log("Handling bomb", data.action);
    setBombs((prev) => {
      prev.push({
        name: data.sender_uuid,
        x: data.action.x,
        y: data.action.y,
        flame: Number(data.message_text),
        exploded: false,
        remove: false,
        fuseTimer: setTimeout(() => explodeBomb(data.action.x, data.action.y), 1900),
        removeTimer: setTimeout(() => removeBomb(), 2700),
      });
      return prev;
    });
  };

  const explodeBomb = (x, y) => {
    setBombs((prev) => {
      prev.map((bomb) => {
        if (bomb.x === x && bomb.y === y) {
          bomb.exploded = true;
        }
      });
      return prev;
    });
  };

  const removeBomb = () => {
    setBombs((prev) => {
      prev.shift();
      return prev;
    });
  };

  function socketSend(type, pos) {
    ws.send({ type: type, action: pos });
  }

  function handleKeys(e) {
    //e.preventDefault();
    if (document.activeElement.id === "messageInput") {
      return;
    }
    if (!player.alive || player.won) return;
    switch (e.code) {
      case "ArrowUp":
        console.log("ArrowUp");
        socketSend("move", { x: player.x, y: limitRange(player.y - 1) });
        break;
      case "ArrowDown":
        console.log("ArrowDown");
        socketSend("move", { x: player.x, y: limitRange(player.y + 1) });
        break;
      case "ArrowLeft":
        console.log("ArrowLeft");
        socketSend("move", { x: limitRange(player.x - 1), y: player.y });
        break;
      case "ArrowRight":
        console.log("ArrowRight");
        socketSend("move", { x: limitRange(player.x + 1), y: player.y });
        break;
      case "Space":
        console.log("Space");
        socketSend("BOMB", { x: player.x, y: player.y });
        break;
      default:
        return;
    }
  }
  keyListener.disconnect();
  keyListener.connect();

  return [
    stats(player),
    map(player, enemies, mapData, bombs),
    (player.name && !player.alive) || player.won
      ? node.div({ class: "overlay" }, player.won ? "You Won!" : "Game Over!")
      : null,
  ];
};

export const initGame = (ws) => {
  init("root", () => Game(ws));
};

export const closeGame = () => {
  //render('root', ['']);
};
