import { node } from '../frawnwok/frawnwok.js';
import { MAP_SIZE } from './game.js';

const map = (player, enemies, mapData, bombs) => {
    let map = [];
    let explosionMap = [...Array(MAP_SIZE)].map(() => [...Array(MAP_SIZE)])
    for (let bomb of bombs) {
        if (bomb.exploded) {
            let right = true
            let up = true
            let left = true
            let down = true
            let target = null
            explosionMap[bomb.y][bomb.x] = ' explosion'
            for (let i = 1; i <= bomb.flame; i++) {

                if (right) {
                    if (bomb.x + i < MAP_SIZE) {
                        target = mapData[bomb.y][bomb.x + i]
                        if (target !== 'wall' && target !== 'box') {
                            explosionMap[bomb.y][bomb.x + i] = " explosion"
                        } else {
                            right = false
                        }
                    } else {
                        right = false
                    }
                }
                if (up) {
                    if (bomb.y - i > 0) {
                        target = mapData[bomb.y - i][bomb.x]
                        if (target !== 'wall' && target !== 'box') {
                            explosionMap[bomb.y - i][bomb.x] = " explosion"
                        } else {
                            up = false
                        }
                    } else {
                        up = false
                    }
                }
                if (left) {
                    if (bomb.x - i > 0) {
                        target = mapData[bomb.y][bomb.x - i]
                        if (target !== 'wall' && target !== 'box') {
                            explosionMap[bomb.y][bomb.x - i] = " explosion"
                        } else {
                            left = false
                        }
                    } else {
                        left = false
                    }
                }
                if (down) {
                    if (bomb.y - i < MAP_SIZE) {
                        target = mapData[bomb.y + i][bomb.x]
                        if (target !== 'wall' && target !== 'box') {
                            explosionMap[bomb.y + i][bomb.x] = " explosion"
                        } else {
                            down = false
                        }
                    } else {
                        down = false
                    }
                }
            }
        } else {
            if (explosionMap[bomb.y][bomb.x]) {
                explosionMap[bomb.y][bomb.x] += ' bomb'
            } else {
                explosionMap[bomb.y][bomb.x] = ' bomb'
            }
        }
    }

    for (let i = 0; i < mapData.length; i++) {
        let row = [];
        for (let j = 0; j < mapData[0].length; j++) {
            let tileClass = `tile ${mapData[i][j]}`;
            if (mapData[i][j].startsWith('player')) {
                tileClass = `tile empty`;
            }

            tileClass += i == player.y && j == player.x ? ' player' + player.id : ''
            for (let k = 0; k < enemies.length; k++) {
                tileClass += i == enemies[k].y && j == enemies[k].x ? ' player' + enemies[k].id : ''
            }
            if (explosionMap[i] && explosionMap[i][j]) {
                tileClass += explosionMap[i][j];
            }

            const tile = node.div({ id: i * MAP_SIZE + j, class: tileClass });
            row.push(tile);
        }
        map.push(node.div({ id: i, class: 'row' }, row));
    }
    return node.div({ id: "map" }, ...map);
};

export default map;
