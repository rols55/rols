import { node } from '../frawnwok/frawnwok.js';

const stats = (player) => {
    if (!player.name) {
        return;
    }
    let lives = []
    for(let i = 0; i < player.life; i++) {
        lives.push(node.div({class: 'life'},''))
    }
    return node.div({ class: 'statsContainer' },
        node.div({ class: 'playerName' }, 'Name: ' + player.name),
        node.div({ class: 'lives' }, ...lives),
        node.div({ class: 'avatar player'+player.id}, ''),
    )
};

export default stats;
