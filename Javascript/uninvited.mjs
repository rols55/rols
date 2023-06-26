import http from 'http';
import { writeFile } from 'node:fs/promises';
const PORT = 5000;
const server = http.createServer(async (req, res) => {
    const { method, url } = req;
    if (method === 'POST') {
        const guest = url.slice(1);
        let body = '';
        req.on('data', (chunk) => {
            body += chunk;
        });
        req.on('end', async () => {
            try {
                await writeFile(`./guests/${guest}.json`, body);
                res.writeHead(201, { 'Content-Type': 'application/json' });
                res.end(body);
            } catch (err) {
                res.writeHead(500, { 'Content-Type': 'application/json' });
                res.end(JSON.stringify({ error: 'server failed' }));
            }
        });
    } else {
        res.writeHead(404, { 'Content-Type': 'application/json' });
        res.end(JSON.stringify({ error: 'not found' }));
    }
});
server.listen(PORT, () => {
    console.log(`Server listening on port ${PORT}`);
});