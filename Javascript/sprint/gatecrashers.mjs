import http from "http";
import fs from "fs";
const users = {
    Caleb_Squires: 'abracadabra',
    Tyrique_Dalton: 'abracadabra',
    Rahima_Young: 'abracadabra'
}
const expected = {
    answer: 'yes',
    drink: 'juice',
    food: 'pizza',
}
const server = http.createServer((req, res) => {
    let { method, url } = req;
    const authHeader = req.headers.authorization
    if (!authHeader || !authHeader.startsWith('Basic ')) {
        res.statusCode = 401;
        res.setHeader('WWW-Authenticate', 'Basic realm="Please enter your username and password"');
        res.end();
        return;
    }
    const authString = authHeader.slice('Basic '.length);
    const [username, password] = Buffer.from(authString, 'base64').toString().split(':');
    if (!users[username] || users[username] !== password) {
        res.statusCode = 401;
        res.setHeader('WWW-Authenticate', 'Basic realm="Please enter your username and password"');
        res.end();
        return;
    }
    let authed = false;
    const fileName = './guests' + url + ".json";
    if (method == "POST") {
        var body = "";
        var parsed;
        req.on("data", (chunk) => {
            body += chunk.toString();
            parsed = JSON.parse(body);
        });
        req.on("end", () => {
            try {
                fs.writeFile(fileName, body, (err) => {
                    if (err) {
                        res.writeHead(500, { "Content-Type": "application/json" });
                        res.end(JSON.stringify({ error: "server failed" }));
                        return;
                    } else {
                        res.writeHead(200, { "Content-Type": "application/json" });
                        res.end(JSON.stringify(expected));
                    }
                });
            } catch (err) {
                res.writeHead(500, { "Content-Type": "application/json" });
                res.end(JSON.stringify({ error: "server failed" }));
            }
        });
    } else {
        res.writeHead(401, { "Content-Type": "application/json" });
        res.end(JSON.stringify({ error: "Authorization Required" }));
    }
});
const port = 5000;
server.listen(port, () => {
    console.log(`Server running on port ${port}`);
});