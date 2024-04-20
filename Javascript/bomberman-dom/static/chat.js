import { useState, node, init } from './frawnwok/chatFramework.js';


const openChat = (ws) => {

    const ENTER = 13
    let chatContainer;
    const [newMessage, setNewMessage] = useState("");
    const [messages, setMessages] = useState([]);

    const sendMessage = (e) => {
        if (e.keyCode === ENTER && !e.shiftKey) {
            e.preventDefault();
            const messageToSend = e.target.value; 
            if (messageToSend.trim() !== '') {
                ws.send({ type: "message", message_text: messageToSend});
            }
            setNewMessage('');
        }
    };
    
    ws.handleOnMessage = (data) => {
        setMessages((prev) => {
            return [...prev, data]
         }); 
       }

    const initializeChatComponents = () => {

        chatContainer = node.div({
            className: 'chatContainer'
        },  node.div({
            className: "chatLogo"
        }), node.div({
            className: "chat"
        },
            node.h2({}, 'Game chat'),
            node.div({
                className: 'messageContainer'
            }, 
                ...messages.slice().reverse().map((message, index) => 
                node.div({
                    key: index,
                    className: 'message'
                },  
                    node.span({
                        className: 'sender'
                    },message.sender_uuid, ":"),
                    node.div({}, message.message_text)
                )
            )
            ),
                node.input({
                    type: 'text',
                    id: 'messageInput',
                    value: newMessage,
                    placeholder: 'Type a message and press enter',
                    className: 'input',
                    onKeyDown: sendMessage,
                }),
        ))
    };



    initializeChatComponents();

    return [chatContainer];
};

export function focusChatInput() {
    const chatInputElement = document.getElementById('messageInput'); 
    if(chatInputElement) chatInputElement.focus();
}

export const initChat = (ws) => {
    init('chat2', () => {
        return openChat(ws);
    });
};

