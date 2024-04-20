import WS from './ws.js'
import {initGame, closeGame} from './game/game.js'
import {initLobby} from './lobby.js'
import { initChat } from './chat.js';


document.addEventListener("DOMContentLoaded", () => {
  // Handle form submission
  document.getElementById("submit").addEventListener("click", () => {
    const name = document.getElementById("nickname").value;

    const ws = new WS(name);

    ws.handleOnOpen = () => {
      console.log("Connected to ws");
      document.getElementById("nickname-container").style.display = "none";
      document.getElementById("logo").style.display = "none";

      //initGame(ws);
      initLobby(ws);
      initChat(ws);
    };

    ws.handleOnClose = () => {
      console.log("Disconnected from ws");
      document.getElementById("nickname-container").style.display = "flex";
      //disconnected from ws so back to "login"

      
      //need to relocate
      closeGame();
    };
  });
});
