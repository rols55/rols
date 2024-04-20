import http from 'http';
import { readFile } from 'node:fs/promises';
const PORT = 5000;
const server = http.createServer(async (req, res) => {
    const { url } = req;
    const guest = url.slice(1);
    try {
        const content = await readFile(`./guests/${guest}.json`, 'utf8');
        const data = JSON.parse(content);
        res.writeHead(200, { 'Content-Type': 'application/json' });
        res.end(JSON.stringify(data));
    } catch (err) {
        if (err.code === 'ENOENT') {
            res.writeHead(404, { 'Content-Type': 'application/json' });
            res.end(JSON.stringify({ error: 'guest not found' }));
        } else {
            res.writeHead(500, { 'Content-Type': 'application/json' });
            res.end(JSON.stringify({ error: 'server failed' }));
        }
    }
});
server.listen(PORT, () => {
    console.log(`Server listening on port ${PORT}`);
});