import { initGame } from "./game/game.js";
import { useState, node, init } from "./frawnwok/frawnwok.js";

const Lobby = (ws) => {
  const [players, setNewPlayer] = useState([]);
  const [timer, setTimer] = useState([]);
  const [timerFinal, setTimerFinal] = useState([11]);

  let LobbyContainer;

  const getTimerMessage = () => {
    if (players.length <= 1) {
      return "You need at least two players to start a game.";
    } else if (players.length >= 2 && players.length < 4 && timerFinal > 10) {
      return `Waiting for more players to join: ${timer} seconds`;
    } else {
      return `Game is starting in  ${timerFinal} seconds`;
    }
  };

  ws.handleOnConnected = (data) => {
    const playerNames = data.players.map((player) => player.name);
    setNewPlayer(() => {
      return playerNames;
    });
  };

  ws.handleOnTimer = (data) => {
    if (data.timerfinal !== undefined) {
      setTimerFinal(data.timerfinal);
      if (data.timerfinal <= 1) {
        initGame(ws); //start the game
      }
    } else if (data.timer !== undefined) {
      setTimer(data.timer);
    }
  };

  const removePlayerPrefix = (nameWithPrefix) => {
    return nameWithPrefix.split("_").pop();
  };

  const initializeLobbyComponents = () => {
    LobbyContainer = node.div(
      {
        className: "lobbyContainer",
      },
      node.div(
        { className: "lobby" },
        node.h2({}, "Lobby"),
        node.div(
          { className: "playerContainer" },
          ...players.map((playerNameWithPrefix, index) => {
            const playerName = removePlayerPrefix(playerNameWithPrefix);
            return node.div(
              { key: index, className: "players" },
              node.span({}, playerName)
            );
          })
        ),
        players.length <= 1
          ? node.div(
              { className: "timer" },
              node.span({}, "You need at least two players to start a game.")
            )
          : node.div({ className: "timer" }, node.span({}, getTimerMessage()))
      )
    );
  };

  initializeLobbyComponents();

  return [LobbyContainer];
};

export const initLobby = (ws) => {
  init("root", () => {
    return Lobby(ws);
  });
};
