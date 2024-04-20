document.addEventListener('DOMContentLoaded', function() {
  loginBtnListener();
  invalidEmailListener();
  signupBtnListener();
  createPostBtnListener();
  clearLoginModalAfterClosing();
  clearSignupModalAfterClosing();
  postBtnListener();
  keepCategoriesStickyTop();
  categoryBtnListener();
  window.onbeforeunload = loadWebSocket();
});

// func to make post text and title change color when hovering on either one
const addHoverEffectToPosts = () => {
  const postTitles = document.querySelectorAll('.post-title');
  const postPreviews = document.querySelectorAll('.card-body');

  postTitles.forEach(function(postTitle) {
      postTitle.addEventListener('mouseover', function() {
          const leftSection = postTitle.parentNode;
          const postHeader = leftSection.parentNode;
          const postText = postHeader.parentNode.querySelector('.card-body');
          postText.classList.add('hovered');
      });
  
      postTitle.addEventListener('mouseout', function() {
          const leftSection = postTitle.parentNode;
          const postHeader = leftSection.parentNode;
          const postText = postHeader.parentNode.querySelector('.card-body');
          postText.classList.remove('hovered');
      });
  });

  postPreviews.forEach(function(postPreview) {
    postPreview.addEventListener('mouseover', function() {
      const postHeader = postPreview.parentNode.querySelector('.card-header');
      const leftSection = postHeader.querySelector('.left-section');
      const postTitle = leftSection.querySelector('.post-title');
      postTitle.classList.add('hovered');
    });

    postPreview.addEventListener('mouseout', function() {
      const postHeader = postPreview.parentNode.querySelector('.card-header');
      const leftSection = postHeader.querySelector('.left-section');
      const postTitle = leftSection.querySelector('.post-title');
      postTitle.classList.remove('hovered');
    });
  });
}

// listen for log in button click
const loginBtnListener = () => {
  const loginForm = document.getElementById("loginForm");
  if (loginForm) {
    loginForm.addEventListener("submit", function(e) {
      e.preventDefault();
      handleLoginSubmit();
    });
  };

  // func to send login data to server and get response
  function handleLoginSubmit() {

    const username = document.getElementById('loginUsername').value;
    const password = document.getElementById('loginPassword').value;

    makePostRequest("/login", "username=" + encodeURIComponent(username) + 
    "&password=" + encodeURIComponent(password), handleLoginResponse);
  }
}

// listen for invalid email address input
const invalidEmailListener = () => {
  let email = document.getElementById("signupEmail");
  email.addEventListener("invalid", function(e) {
    e.preventDefault();
    if (!email.validity.valid) {
      const errMsgField = document.getElementById("signupErrorMsg");
      errMsgField.innerHTML = "E-mail address is not valid";
      errMsgField.removeAttribute('hidden');
      document.getElementById('signupPassword').value = "";
      document.getElementById('signupPassword2').value = "";
      document.getElementById('signupEmail').value = "";
    }
  });
}

// listen for sign up button click
const signupBtnListener = () => {
  const signupForm = document.getElementById("signupForm");
  signupForm.addEventListener("submit", function(e) {
    e.preventDefault();
    handleSignupSubmit();
  });

  // func to send signup data to server and get response
  function handleSignupSubmit() {
    const username = document.getElementById('signupUsername').value;
    const age = document.getElementById('birthday').value;
    const gender = document.getElementById('signupSex').value;
    const firstName = document.getElementById('signupfirstname').value;
    const lastName = document.getElementById('signupLastname').value;
    const email = document.getElementById('signupEmail').value;
    const password = document.getElementById('signupPassword').value;
    const password2 = document.getElementById('signupPassword2').value;

    if (password !== password2) {
      const errMsgField = document.getElementById("signupErrorMsg");
      errMsgField.innerHTML = "Passwords don't match";
      errMsgField.removeAttribute('hidden');
      document.getElementById('signupPassword').value = "";
      document.getElementById('signupPassword2').value = "";
      return
    };

    makePostRequest("/register", 
    "&age=" + encodeURIComponent(age) +
    "&email=" + encodeURIComponent(email) +
    "&gender=" + encodeURIComponent(gender) +
    "&firstName=" + encodeURIComponent(firstName) +
    "&lastName=" + encodeURIComponent(lastName) +
    "&username=" + encodeURIComponent(username) + 
    "&password=" + encodeURIComponent(password),
    handleSignupResponse)
  }
}

// listen for create post button click
const createPostBtnListener = () => {
  const createPostForm = document.getElementById("createPostForm");
  createPostForm.addEventListener("submit", function(e) {
    e.preventDefault();
    submitCreatePost();
  });

  // func to send new post data to server and get response
  function submitCreatePost() {
    const category = Array.from(document.getElementById('newPostCategory').selectedOptions).map(({ value }) => value);
    const title = document.getElementById('newPostTitle').value;
    const text = document.getElementById('newPostText').value;

    cleanCreatePostFormErrors()
  
    makePostRequest("/post/create", 
    "category=" + encodeURIComponent(category) +
    "&title=" + encodeURIComponent(title) + 
    "&text=" + encodeURIComponent(text),
    handleCreatePostResponse)
  }
}

//General function for post requests
function makePostRequest(url, data, callback) {
  const xhr = new XMLHttpRequest();
  xhr.open("POST", url, true);
  xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");

  xhr.onreadystatechange = function() {
    if (this.readyState === XMLHttpRequest.DONE && this.status === 200) {
      callback(JSON.parse(this.responseText));
    }
  }
  
  xhr.send(data);
}

function handleLoginResponse(response) {
  if (response.status === 'OK') {
    const loginModal = bootstrap.Modal.getInstance(document.getElementById('loginModal'));
    loginModal.hide();
    loadUserIndexPage(response);
    loadWebSocket();
  } else {
    // Display an error message
    const loginMsgField = document.getElementById("loginMsgField");
    loginMsgField.innerHTML = response.msg;
    if (loginMsgField.classList.contains("text-success")) {
      loginMsgField.classList.remove("text-success");
      loginMsgField.classList.add("text-danger");
    }
    loginMsgField.removeAttribute('hidden');
    document.getElementById('loginPassword').value = "";
    if (response.msg !== "No password entered" && 
        response.msg !== "Wrong password") {
      document.getElementById('loginUsername').value = "";
    }
  }
};

function loadUserIndexPage(resp) {
  const nav =  document.querySelector(".navbar-nav");
  nav.removeAttribute("hidden");
  nav.querySelector(".navbar-text").innerHTML = `Welcome, ${resp.user.username}!`;

  const postsHTML = `
      <div class="col-lg-2 pt-5 ps-0 pe-0">
        <div id="usersWrapper" class="sticky-top">
            <h5>Chat</h5>
            <div id="chatbox" class="d-flex flex-wrap align-items-start">
                <div id="userPanel"></div>
            </div>
        </div>
      </div>
      <div class="col-lg-8 pt-4 ps-0">
        <h3 class="currentCatTitle fw-bold fs-4">All</h3>
        <div id="postsList">
            ${(function () {
              let result = "";
              resp.posts.forEach((post) => {
                result += `
                <button id="postButton" class="btn-light me-2 bg-transparent border-0" data-bs-toggle="modal" data-bs-target="#postModal" data-id="${post.id}" data-category="${post.category_ids}">
                  <div class="card mb-3 border-0 pb-3">
                      <div class="card-header d-flex bg-transparent border-0 ps-0 pb-0 justify-content-between"> 
                          <div class="left-section d-flex align-items-center">
                              <a class="post-title fw-bolder fs-5 text-decoration-none me-2" href="/post/${post.id}">
                                  ${post.title}</a>
                          </div>
                          <div class="right-section d-flex flex-nowrap align-items-center">
                              <span class="username me-3">${post.author}</span>
                              <i class="fa-regular fa-comment fa-xl me-1"></i>
                              <span class="comment-count post_${post.id}">${post.comments_qty}</span>
                          </div>
                      </div>
                      <a class="card-body text-decoration-none ps-0 pt-1" href="/post/${post.id}">
                          <p class="post-text mb-0">
                              ${post.text}
                          </p>
                      </a>
                  </div>
                </button>
                `
              });
              return result
            })()}
        </div>
      </div>
      <div class="col-lg-2 ps-0 pe-0">
          <div id="categoryWrapper" class="sticky-top">
              <h5>Topics</h5>
              <div id="categories" class="d-flex flex-wrap align-items-start">
                  <a id="0" href="#" class="categoryBubble">All categories</a>
                  ${(function () {
                    let result = "";
                    resp.categories.forEach( c => {
                      result += `<a id="${c.id}" href="#" class="categoryBubble">${c.category}</a>`
                    });
                    return result;
                  })()}
              </div>
          </div>
      </div>
`
  const elm = document.querySelector(".container .row")
  elm.innerHTML = postsHTML
  postBtnListener();
  keepCategoriesStickyTop();
  categoryBtnListener();
}

function categoryBtnListener() {
  const categories = document.querySelectorAll("#categories .categoryBubble");
  categories.forEach(btn => {
    btn.addEventListener("click", function(e) {
      e.preventDefault();
      currentCategory = document.querySelector(".currentCatTitle");
      currentCategory.innerHTML = (btn.id === "0" ? "All" : btn.innerHTML);
      postBtns = document.querySelectorAll("#postButton")
      postBtns.forEach(post => {
        var ids = post.dataset.category.split(",")
        if (ids.includes(btn.id) || btn.id === "0") {
          post.removeAttribute("hidden");
        } else {
          post.setAttribute("hidden", "true");
        }
      });
    });
  })
};


function handleSignupResponse(response) {
  if (response.status === 'OK') {
    // Close the signup modal and open login one
    const signupModal = bootstrap.Modal.getInstance(document.getElementById('signupModal'));
    signupModal.hide();
    const loginModalElement = document.getElementById('loginModal');
    let loginModal = bootstrap.Modal.getInstance(loginModalElement);
    if (!loginModal) {
        loginModal = new bootstrap.Modal(loginModalElement);
    }
    loginModal.show();
    const loginMsgField = document.getElementById("loginMsgField");
    loginMsgField.innerHTML = "Account created, log in to continue";
    loginMsgField.classList.remove("text-danger");
    loginMsgField.classList.add("text-success");
    loginMsgField.removeAttribute('hidden');
  } else {
    // Display an error message
    const errMsgField = document.getElementById("signupErrorMsg");
    errMsgField.innerHTML = response.msg;
    errMsgField.removeAttribute('hidden');
    document.getElementById('signupPassword').value = "";
    document.getElementById('signupPassword2').value = "";
    if (response.msg.startsWith("Username")) {
      document.getElementById('signupUsername').value = "";
    } else if (response.msg.startsWith("Account with")) { // For emails which are already used for account
      document.getElementById('signupEmail').value = "";
    }
  }
}

function handleCreatePostResponse(response) {
  if (response.status === 'OK') {
    // Close the modal and reload the page
    const createPostModal = bootstrap.Modal.getInstance(document.getElementById('createPostModal'));
    createPostModal.hide();

    const postsHTML =`
<button id="postButton" class="btn-light me-2 bg-transparent border-0" data-bs-toggle="modal" data-bs-target="#postModal" data-id="${response.post.id}" data-category="${response.post.category_ids}">
  <div class="card mb-3 border-0 pb-3">
      <div class="card-header d-flex bg-transparent border-0 ps-0 pb-0 justify-content-between"> 
          <div class="left-section d-flex align-items-center">
              <a class="post-title fw-bolder fs-5 text-decoration-none me-2" href="/post/${response.post.id}">
                  ${response.post.title}</a>
          </div>
          <div class="right-section d-flex flex-nowrap align-items-center">
              <span class="username me-3">${response.post.author}</span>
              <i class="fa-regular fa-comment fa-xl me-1"></i>
              <span class="comment-count post_${response.post.id}">${response.post.comments_qty}</span>
          </div>
      </div>
      <a class="card-body text-decoration-none ps-0 pt-1" href="/post/${response.post.id}">
          <p class="post-text mb-0">
              ${response.post.text}
          </p>
      </a>
  </div>
</button>
`
    const postsList = document.getElementById('postsList');
    postsList.innerHTML = postsHTML + postsList.innerHTML;
    postBtnListener();
  } else {
    // Display an error message
    const errorMsg = response.msg;
    if (errorMsg === "Category missing") {
      document.getElementById("newPostCategory").classList.add("is-invalid")
      document.getElementById("newPostCategoryLabel").innerHTML = "Pick a category which describes your post the most";
    } else if (errorMsg === "Title missing") {
      document.getElementById("newPostTitle").classList.add("is-invalid")
      document.getElementById("newPostTitleLabel").innerHTML = "Your post needs a title";
    } else if (errorMsg === "Text missing") {
      document.getElementById("newPostText").classList.add("is-invalid")
      document.getElementById("newPostTextLabel").innerHTML = "Add some text to your post";
    } else if (errorMsg === "Session expired") {
      alert("Your session has expired")
    } else if (errorMsg === "Title too long"){
      document.getElementById("newPostTitle").classList.add("is-invalid")
      document.getElementById("newPostTitleLabel").innerHTML = "Your post's title is too long, max 60 characters";
    }
  }
}

// clear form fields and remove error message if login modal is closed
const clearLoginModalAfterClosing = () => {
  const loginModal = document.getElementById('loginModal');
  if (loginModal) {
    loginModal.addEventListener('hidden.bs.modal', function () {
        document.getElementById('loginForm').reset();
        const loginMsgField = document.getElementById('loginMsgField');
        loginMsgField.setAttribute('hidden', '');
        if (loginMsgField.classList.contains("text-success")) {
          loginMsgField.classList.remove("text-success");
          loginMsgField.classList.add("text-danger");
        }
    });
  }
}

// clear form fields and remove error message if signup modal is closed
const clearSignupModalAfterClosing = () => {
  const signupModal = document.getElementById('signupModal');
  signupModal.addEventListener('hidden.bs.modal', function () {
      document.getElementById('signupForm').reset();
      document.getElementById('signupErrorMsg').setAttribute('hidden', '');
  });

  // clear form fields and remove error message if create post modal is closed
  const createPostModal = document.getElementById('createPostModal');
  createPostModal.addEventListener('hidden.bs.modal', function () {
    document.getElementById('createPostForm').reset();
    cleanCreatePostFormErrors()
  });
}

//Function to clean error messages from create post form (But keeps input values)
function cleanCreatePostFormErrors() {
  document.getElementById("newPostCategory").classList.remove("is-invalid")
  document.getElementById("newPostCategoryLabel").innerHTML = "Pick a category";
  document.getElementById("newPostTitle").classList.remove("is-invalid")
  document.getElementById("newPostTitleLabel").innerHTML = "Post title";
  document.getElementById("newPostText").classList.remove("is-invalid")
  document.getElementById("newPostTextLabel").innerHTML = "Post text";
}

//Keep categories and users from hiding when scrolling
const keepCategoriesStickyTop = () => {
  const navbar = document.querySelector('.navbar');
  const navbarHeight = navbar.offsetHeight;
  const categoryWrapper = document.querySelector('#categoryWrapper');
  if (categoryWrapper) {
    categoryWrapper.style.top = navbarHeight+40 + 'px';
  }
  const usersWrapper = document.querySelector('#usersWrapper');
  if (usersWrapper) {
    usersWrapper.style.top = navbarHeight+40 + 'px';
  }
}

// Helper function to swap classes
function swapClasses(element, remove, add) {
  element.classList.remove(remove);
  element.classList.add(add);
}

//Set reaction button default state and hovering behavior
const setReactionBtnBehavior = (reactionBtn) => {
  const userSelected = reactionBtn.dataset.selected === "true";
  const toggleClasses = userSelected ? ["fa-regular", "fa-solid"] : ["fa-solid", "fa-regular"];

  // Set the initial state
  swapClasses(reactionBtn, toggleClasses[0], toggleClasses[1]);
  reactionBtn.style.cursor = "pointer";

  // Toggle on hover
  reactionBtn.addEventListener("mouseover", function() {
    swapClasses(reactionBtn, toggleClasses[1], toggleClasses[0]);
  });

  // Toggle again on mouse leave
  reactionBtn.addEventListener("mouseleave", function() {
    swapClasses(reactionBtn, toggleClasses[0], toggleClasses[1]);
  });
}

//Change state of reaction
const sendReactionRequest = (reactionBtn) => {
  let userId = document.getElementById("userId").getAttribute("value");
  let postId = reactionBtn.dataset.id;

  // Create reaction data
  let reactionData = {
    userID: parseInt(userId, 10),
    postID: parseInt(postId, 10), 
    reaction: reactionBtn.classList.contains("giveLike") ? true : false,
    addOrRemove: (reactionBtn.dataset.selected === "true") ? false : true,
    isPost: (reactionBtn.dataset.comment === "true") ? false : true
  };

  // Post the reaction
  return fetch('/reactions', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(reactionData)
  })
  .then(response => {
    if (!response.ok) {
      return response.json().then(error => alert(error.error));
    }
    return response.json();
  })
  .then(data => {
    // Show new likes/dislikes amount next to reaction button
    reactionBtn.nextElementSibling.innerHTML = data.reactionAmount;
    if (reactionData.addOrRemove) {
      reactionBtn.dataset.selected = "true";
    } else {
      reactionBtn.dataset.selected = "false";
    }
    
    setReactionBtnBehavior(reactionBtn);
  })
  .catch(error => console.error('Error:', error));
}

// Comments

//expand new comment form on click
const expandCommentFormListener = () => {
  const createCommentTextbox = document.getElementById("newCommentText");
  if (createCommentTextbox != null) {
    createCommentTextbox.addEventListener("click", function() {
      expandCommentForm();
    });

    function expandCommentForm() {
      const newCommentTextbox = document.getElementById("newCommentText");
      const commentFormButtons = document.getElementById("commentFormButtons");
      newCommentTextbox.style.height = "20vh";
      commentFormButtons.classList.remove("d-none");
    }
  }
}

//close new comment form when clicking Cancel
const cancelCommentFormListener = () => {
  const cancelNewComment = document.getElementById("cancelCreateComment");
  if (cancelNewComment != null) {
    cancelNewComment.addEventListener("click", function() {
      closeCommentForm();
    });

    function closeCommentForm() {
      const newCommentTextbox = document.getElementById("newCommentText");
      const commentFormButtons = document.getElementById("commentFormButtons");
      newCommentTextbox.style.height = "calc(3.5rem + 2px)";
      commentFormButtons.classList.add("d-none");
    }
  }
}

// listen for create comment button click
const createCommentBtnListener = () => {
  const createCommentForm = document.getElementById("createCommentForm");
  if (createCommentForm != null) {
    createCommentForm.addEventListener("submit", function(e) {
      e.preventDefault();
      submitCreateCommentForm();
    });

    // func to send new comment data to server and get response
    function submitCreateCommentForm() {
      const postId = document.getElementById('newCommentPostId').value;
      const text = document.getElementById('newCommentText').value;

      document.getElementById("newCommentText").classList.remove("is-invalid")
      document.getElementById("newCommentTextLabel").innerHTML = "Comment text";

      makePostRequest("/comment/create", 
      "postid="+ encodeURIComponent(postId) +
      "&text=" + encodeURIComponent(text),
      handleCreateCommentResponse)
    }
  }
}

function handleCreateCommentResponse(response) {
  if (response.status === 'OK') {
    document.getElementById('newCommentText').value = "";
    const postModal = document.getElementById("postModalContainer")
    postModal.insertAdjacentHTML("beforeend", getCommentsHTML([response.comment], response.user ? response.user.id : null));
    let count = document.querySelector(".card .comment-count")
    count.innerHTML = Number(count.innerHTML) + 1
    count = document.querySelector(".comment-count.post_"+response.comment.post_id)
    count.innerHTML = Number(count.innerHTML) + 1
    const form = document.getElementById("delete_comment_"+response.comment.id)
    form.addEventListener("submit", function(e) {
      e.preventDefault();
      const id = e.submitter.dataset.id;
      makePostRequest(form.action, 
      "id="+ encodeURIComponent(id), response => {
        if (response.status === 'OK') {
          let comment = document.querySelector("#comment_"+id)
          comment.remove()
          let count = document.querySelector(".card .comment-count")
          count.innerHTML -= 1
          count = document.querySelector(".comment-count.post_"+response.post_id)
          count.innerHTML -= 1
        } else {
          alert(response.msg)
        }
      })
    });
  } else {
    // Display an error message
    const errorMsg = response.msg;
    if (errorMsg === "Text missing") {
      document.getElementById("newCommentText").classList.add("is-invalid")
      document.getElementById("newCommentTextLabel").innerHTML = "Add some text to your comment";
    } else if (errorMsg === "Session expired") {
      alert("Your session has expired")
    } 
  }
}

function toggleBold() {
  const links = document.getElementById("profileLinks").children
  var url = window.location.href;
  for (const lnk of links) {

    if (url === lnk.href) {
      lnk.classList.add("fw-semibold");
    } else {
      lnk.classList.remove("fw-semibold");
    }
  }
}

function deletePostBtnListener(){
  const deleteForm = document.querySelectorAll("#modifyPostForm");
  if (deleteForm != null) {
    for (const form of deleteForm){
      form.addEventListener("submit", function(e) {
        e.preventDefault();
        const id = e.submitter.dataset.id;
        makePostRequest(form.action, 
        "id="+ encodeURIComponent(id), response => {
          if (response.status === 'OK') {
            const postModal = bootstrap.Modal.getInstance(document.getElementById('postModal'));
            if (postModal) {
              postModal.hide();
            }
            const postBtn = document.querySelector(`button[data-id="${id}"]`);
            if (postBtn) {
              postBtn.remove();
            }
            //location.replace(response.redirect);
          } else {
            alert(response.msg)
          }
        })
      });
    }
  }
}

// Post Modal
const postBtnListener = () => {
  const modalButtons = document.querySelectorAll("#postButton");
  modalButtons.forEach(function(btn) {
    btn.addEventListener("click", function(e) {
      e.preventDefault();
      const id = btn.dataset.id
      makeGetRequest("/post/"+id, function(resp) {
        const postModal = document.getElementById("postModalContainer")
        postModal.innerHTML = getPostHTML(resp);
        
        expandCommentFormListener();
        cancelCommentFormListener();
        createCommentBtnListener();
        deleteCommentBtnListener();
        deletePostBtnListener();
      })
    });
  });
}

function deleteCommentBtnListener(){
  const deleteForm = document.querySelectorAll(".modifyCommentForm");
  if (deleteForm != null) {
    for (const form of deleteForm){
      form.addEventListener("submit", function(e) {
        e.preventDefault();
        const id = e.submitter.dataset.id;
        makePostRequest(form.action, 
        "id="+ encodeURIComponent(id), response => {
          if (response.status === 'OK') {
            let comment = document.querySelector("#comment_"+id)
            comment.remove()
            let count = document.querySelector(".card .comment-count")
            count.innerHTML -= 1
            count = document.querySelector(".comment-count.post_"+response.post_id)
            count.innerHTML -= 1
          } else {
            alert(response.msg)
          }
        })
      });
    }
  }
}

function makeGetRequest(url, callback) {
  const xhr = new XMLHttpRequest();
  xhr.open("GET", url, true);
  xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
  xhr.onreadystatechange = function() {
    if (this.readyState === XMLHttpRequest.DONE && this.status === 200) {
      callback(JSON.parse(this.responseText));
    }
  }
  xhr.send();
}



function getPostHTML(response) {
  return `
  <div class="card mb-3 border-0 pb-3">
    <div class="card-header d-flex bg-transparent border-0 ps-0 pb-0 justify-content-between">
      <div class="left-section d-flex align-items-center">
        <span class="post-title fw-bolder fs-5 text-decoration-none me-2">
          ${response.post.title}</span>
      </div>
      <div class="right-section d-flex flex-nowrap align-items-center">
        <span class="username me-3">${response.post.author}</span>
        <i class="fa-regular fa-comment fa-xl me-1"></i>
        <span class="comment-count">${response.post.comments_qty}</span>
      </div>
    </div>
    <!-- Text -->
    <div class="card-body text-decoration-none ps-0 pt-1">
      <p class="post-text mb-0">
        ${response.post.text}
      </p>
    </div>
    <div class="mt-3">
    <form id="modifyPostForm" method="post" action="/post/delete" ${response.post.user_id === response.user.id ? "" : "hidden"}>
        <input type="submit" id="deletePostBtn" class="btn btn-primary" data-id="${response.post.id}" value="Delete"/>
    </form>
</div>
  </div>
  ${ response.user ? getAddCommentFormHTML(response.post) : ""}
  ${ response.comments ? getCommentsHTML(response.comments, response.user ? response.user.id : null) : ""}
`
}

function getAddCommentFormHTML(post) {
  return `
<form id="createCommentForm" method="post">
  <div class="form-floating mb-3">
      <textarea id="newCommentText" class="form-control" placeholder=" "></textarea>
      <label id="newCommentTextLabel" class="text-secondary" for="newCommentTextLabel">Add new comment...</label>
  </div>
  <div id="commentFormButtons" class="form-group text-center mb-4 d-none">
      <input type="button" id="cancelCreateComment" class="btn btn-secondary" value="Cancel"/>
      <input type="submit" id="submitCreateCommentForm" class="btn btn-primary" value="Create Comment"/>
  </div>
  <input id="newCommentPostId" value="${post.id}" hidden/>
</form>
`
}

function getCommentsHTML(comments, user_id=null) {
  let comHTML = ""
  for (const com of comments) {
    comHTML += `
    <div class="comment-box card mb-3 p-3 pt-1 border-0" id="comment_${com.id}">
      <div class="card-header d-flex bg-transparent border-0 ps-0 pb-1 pe-0 justify-content-between">
          <div class="left-section d-flex align-items-center">
              <span class="fw-semibold fs-7">${com.author}:</span>
          </div>
      </div>
      <p class="mb-0">${com.text}</p>
      
      ${user_id && user_id === com.user_id ?
      `<div class="mt-3">
          <form id="delete_comment_${com.id}" class="modifyCommentForm" method="post"  action="/comment/delete">
                  <input type="submit" id="deleteCommentBtn" class="btn btn-primary" data-id="${com.id}" value="Delete"/>
          </form>
      </div>`
      : ""}
    </div>
    `
  }
  return comHTML
}

let offset = 0
async function fetchMessages(target, offset) {
  // Send a request for message history by user UUID
    //uuid.value
    fetch("/history" + "?target=" + target + "&offset=" + offset)
    .then(async response => {
      if (response.ok) {
        data = await response.json();
        prependLog(data)
      } else {
        alert(error.msg);
      }
    })
    .catch(error => {
      messages.removeEventListener("wheel", wheelUp);
      messages.removeEventListener("scroll", scrolledToTop);
    }); 
}

function throttle(mainFunction, delay) {
  let timer = null;
  return async (...args) => {
    if (timer === null) {
      mainFunction(...args);
      timer = setTimeout(() => {
        timer = null;
      }, delay);
    } else {
      offset -= 1;
    }
  };
}

const throttledMessages = throttle(fetchMessages, 1000);


let messages = document.querySelector('#messages');
let target
async function openChat(target_id, username) {
  let messages = document.getElementById('messages');
  messages.removeEventListener("wheel", wheelUp);
  messages.removeEventListener("scroll", scrolledToTop)
  clearChat()
  target = target_id
  await throttledMessages(target, 0);
  messages.addEventListener("wheel", wheelUp)
  messages.addEventListener("scroll", scrolledToTop)

  const inputUUID = document.getElementById('uuid');
  if (inputUUID) {
    inputUUID.value = target_id;
  }
  const inputUsername = document.getElementById('username');
  if (inputUsername) {
    inputUsername.value = username;
  }
  const titleUsername = document.getElementById('titleUsername');
  if (titleUsername) {
    titleUsername.innerHTML = username;
  }
}

async function wheelUp(evt) {
  if (evt.deltaY < 0) {
    offset += 1
    await throttledMessages(target, offset);
  }
}

async function scrolledToTop() {
  if (messages.scrollTop === 0) {
    offset += 1
    await throttledMessages(target, offset);
  }
}

function clearChat(){
  offset = 0
  messages = document.getElementById('messages');
  const divs = messages.querySelectorAll("div");
  divs.forEach(function(div) {
    div.remove()
  })
}

const userMap = new Map();

async function getUsers() {
  fetch("/users")
  .then(async response => {
    /*
    if (!response.ok) {
      return response.json().then(error => alert(error.msg));
    }
    */
    return response.json();
  })
  .then(data => {
    const userPanel = document.querySelector("#userPanel");
    if (userPanel) {
      userPanel.innerHTML = '<ul id="userList"></ul>'
      const list = userPanel.querySelector('#userList')
      data.users.forEach(user => {
        userMap.set(user.uuid, user.username);
        const listItem = document.createElement('li');
        listItem.id = user.uuid;
        listItem.className = 'userBubble';
        listItem.innerHTML = `<span class="online-indicator"></span>${user.username}`;
        listItem.setAttribute('data-bs-toggle','modal');
        listItem.setAttribute('data-bs-target', '#chat');
        list.appendChild(listItem);
        listItem.addEventListener("click", () => {
          openChat(user.uuid, user.username)
        })
      });
      if (conn) {
        let get_online = {
          type : "get_online",
        }
        conn.send(JSON.stringify(get_online));
      }
    }

  })
  .catch(error => console.error('Error:', error));
  return true;
}

function divConstructor(data){
  var item = document.createElement("div");
  let datetime = new Date(data.timestamp);
  let date = `${datetime.getDate() < 10 ? '0' : ''}${datetime.getDate()}/${datetime.getMonth()+1 < 10 ? '0': ''}${datetime.getMonth()+1}/${datetime.getFullYear()}`;
  let time = `${datetime.getHours() < 10 ? '0' : ''}${datetime.getHours()}:${datetime.getMinutes() < 10 ? '0' : ''}${datetime.getMinutes()}`;
  datetime =  `${date} ${time}`;
  if (uuid.value !== data.reciver_uuid) {
    const username = userMap.get(data.sender_uuid);
    item.style.textAlign = "left";
    var infoElementUser = document.createElement("span");
    infoElementUser.className = "userMessage"; 
    infoElementUser.innerText =`${username} ${datetime}\n`

    var messageTextElementUser = document.createElement("span");
    messageTextElementUser.className = "userMessage";
    messageTextElementUser.innerText = `${data.message_text}`;
    item.appendChild(infoElementUser);
    item.appendChild(messageTextElementUser);
  } else {
    item.style.textAlign = "right";
    var infoElement = document.createElement("span");
    infoElement.className = "myMessage"; 
    infoElement.innerText =`Me ${datetime}:\n`

    var messageTextElement = document.createElement("span");
    messageTextElement.className = "myMessage"; 
    messageTextElement.innerText = `${data.message_text}`;
    item.appendChild(infoElement);
    item.appendChild(messageTextElement);
  }
  return item;
}

function prependLog(data) {
  // Get the current scroll position
  const scrollPosition = messages.scrollTop;
  // Calculate the height of the container with new content
  const containerHeight = messages.scrollHeight;
  data.forEach(message =>
    messages.prepend(divConstructor(message))
  );
   // Set the scroll position back
   messages.scrollTop = scrollPosition;
   if (messages.scrollHeight !== containerHeight) {
    messages.scrollTop = scrollPosition + (messages.scrollHeight - containerHeight);
  }
}

function appendLog(data) {
  const item = divConstructor(data);
  var doScroll = messages.scrollTop > messages.scrollHeight - messages.clientHeight - 1;
  messages.appendChild(item);
  if (doScroll) {
      messages.scrollTop = messages.scrollHeight - messages.clientHeight;
  }
}

var conn;
function loadWebSocket() {
    var msg = document.getElementById("msg");
    document.getElementById("chatInputForm").onsubmit = function () {
        if (!conn) {
            return false;
        }
        if (!msg.value || !uuid.value) {
            return false;
        }
        let message = {
          type: "message",
          reciver_uuid: uuid.value,
          message_text: msg.value
        };
        conn.send(JSON.stringify(message));
        getUsers();
        msg.value = "";
        return false;
    };

    if (window["WebSocket"]) {
        conn = new WebSocket("ws://" + document.location.host + "/ws");
        conn.onclose = function (evt) {
            console.log("WebSocket connection closed");
        };
        conn.onerror =  evt => console.log(evt);
        conn.onmessage = function (evt) {
            var data = JSON.parse(evt.data)
            switch (data.type) {
              case "message":
                appendLog(data);
                getUsers();
                break;
              case "connected":
                getUsers();
                updateUsers([data.sender_uuid], true);
                break;
              case "disconnected":
                updateUsers([data.sender_uuid], false);
                break;
              case "get_online":
                updateUsers(data.online, true);
                break;
            }
        };
        conn.onopen = function (evt) {
          console.log("WebSocket connection established");
          getUsers();
        }
    } else {
        alert("Your browser does not support WebSockets.")
    }
};

function updateUsers(users, online) {
  users.forEach( uuid => {
    const listItem = document.getElementById(`${uuid}`);
    if (listItem) {
      const indicator = listItem.querySelector(".online-indicator");
      if (online) {
        indicator.classList.add('online');  
      } else {
        indicator.classList.remove('online');
      }
    }
  })
}