function hasSessionCookie() {
    const cookies = document.cookie.split('; ');
    return cookies.some(cookie => cookie.startsWith("sessionID"));
}

function getUsername() {
    const username = document.cookie.split(', ')[1]
    return username
}

const usernameHeader = document.getElementById("usernameHeader");

if (hasSessionCookie()) {
    document.getElementById('login').style.display = 'none'; // Show button 1
    document.getElementById('logout').style.display = 'block'; // Show button 2
    document.getElementById("commentSubmissionForm").style.display = "block";
    document.getElementById("createPost").style.display = "block";
    usernameHeader.textContent = getUsername();
    usernameHeader.style.display = "block";
} else {
    document.getElementById('logout').style.display = 'none'; // Show button 2
    document.getElementById('login').style.display = 'block'; // Show button 1
    document.getElementById('createPost').style.display = 'none';
    usernameHeader.style.display = "none";
    document.getElementById("commentSubmissionForm").style.display = "none";

}

document.getElementById("createPost").onclick = function () {
    location.href = "/createPost";
};

const postTags = document.querySelector(".post-tags")
let tags = postTags.dataset.categories.split(",")

for (let tag of tags) {
    let tagElement = document.createElement("li");
    tagElement.textContent = tag;
    postTags.appendChild(tagElement);
}