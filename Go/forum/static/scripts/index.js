// Define some global variables for convenience
// Select the template element
const postTemplate = document.getElementById("post-template");
// Select the posts section
const postsSection = document.getElementById("posts");
// Select the filter form
const filterForm = document.getElementById("filter-form");
// Select tag box from filter section
const tagBox = document.getElementById("tags")
// Select the sort by select element
const sortBySelect = document.getElementById("sort-by");
// Define an array of objects to store the posts data
let posts = [];

const usernameHeader = document.getElementById("usernameHeader");

function getUsername() {
  const username = document.cookie.split(', ')[1]
  return username
}

document.addEventListener("DOMContentLoaded", function () {
  if (hasSessionCookie()) {
    document.getElementById('login').style.display = 'none'; // Show button 1
    document.getElementById('logout').style.display = 'block'; // Show button 2
    document.getElementById('createPost').style.display = 'block';
    document.getElementById("myPosts").style.display = "block";
    usernameHeader.textContent = getUsername();
  } else {
    document.getElementById('logout').style.display = 'none'; // Show button 2
    document.getElementById('login').style.display = 'block'; // Show button 1
    document.getElementById('createPost').style.display = 'none';
    document.getElementById("myPosts").style.display = "none";
    sortBySelect.style.display = "none";
    usernameHeader.textContent = "";
  }
});

document.getElementById("createPost").onclick = function () {
  location.href = "/createPost";
};
// Define a function to create a post element from the template and the data
function createPostElement(post) {
  // Clone the template content
  let postElement = postTemplate.content.cloneNode(true);
  // Select the elements inside the template
  let postTitle = postElement.querySelector(".post-title");
  let postAuthor = postElement.querySelector(".post-author");
  let postText = postElement.querySelector(".post-text");
  let postTags = postElement.querySelector(".post-tags");
  let likeIcon = postElement.querySelector(".like-icon");
  let likeCount = postElement.querySelector(".like-count");
  let dislikeIcon = postElement.querySelector(".dislike-icon");
  let dislikeCount = postElement.querySelector(".dislike-count");
  let postId = postElement.getElementById("post-id");
  let likePost = postElement.getElementById("likePost")
  let dislikePost = postElement.getElementById("dislikePost")
  let parent = postElement.querySelector(".post")

  parent.dataset.date = post.Id;
  parent.dataset.likes = post.Likes;
  parent.dataset.dislikes = post.Dislikes;

  likePost.value = post.Id;
  dislikePost.value = post.Id;
  postId.value = post.Id;
  // Populate the elements with the data
  postTitle.textContent = post.Title;
  postAuthor.textContent = `by ${post.Username}`;
  postText.textContent = post.Body;
  // Loop through the tags and create li elements for each one
  let tags = post.Categories.split(",")
  for (let tag of tags) {
    let tagElement = document.createElement("li");
    tagElement.textContent = tag;
    postTags.appendChild(tagElement);
    appendTagToFilterBox(tag);
  }
  // Set the initial values for the likes and dislikes
  likeIcon.textContent = "👍";
  likeCount.textContent = post.Likes;
  dislikeIcon.textContent = "👎";
  dislikeCount.textContent = post.Dislikes;
  // Return the post element
  return postElement;
}

function appendTagToFilterBox(tag) {
  let checkboxContainer = document.createElement("div")
  checkboxContainer.className = "checkbox-container"
  let checkbox = document.createElement("input")
  checkbox.type = "checkbox"
  checkbox.name = `${tag}`
  checkboxContainer.appendChild(checkbox)
  let tagLabel = document.createElement("label")
  tagLabel.htmlFor = `${tag}`
  tagLabel.textContent = `${tag}`
  checkboxContainer.appendChild(tagLabel)
  tagBox.appendChild(checkboxContainer)
}

function fetchPosts() {
  fetch('/morePosts')
    .then(response => {
      if (!response.ok) {
        throw new Error('Could not fetch posts');
      }
      return response.json();
    })
    .then(data => {
      renderPosts(data);
    })
    .catch(error => {
      // Handle the error
      console.error(error);
    });
}

// Define a function to render the posts on the page
function renderPosts(data) {
  tagBox.innerHTML = "";
  postsSection.innerHTML = "";
  // Loop through the posts array
  for (let post of data) {
    // Create a post element for each post object
    let postElement = createPostElement(post);
    // Append the post element to the posts section
    postsSection.appendChild(postElement);
  }
  posts = Array.from(document.querySelectorAll(".post"))
}

function renderFilteredPosts(posts) {
  postsSection.innerHTML = "";

  for (let post of posts) {
    postsSection.appendChild(post)
  }
}

// Define a function to filter the posts by the selected tags
function filterPostsByTags(posts) {
  // Get the checked tags from the filter form
  let checkedTags = filterForm.querySelectorAll("input[type=checkbox]:checked");
  // Convert the checked tags to an array of tag names
  let tagNames = Array.from(checkedTags).map(tag => tag.name);
  // Check if the tag names array is not empty
  let match = false;

  if (tagNames.length > 0) {
    // Filter the posts that have at least one of the tag names
    let filteredPosts = posts.filter(post => {
      let category = post.querySelector(".post-tags").outerText
      if (category.includes("\n")) {
        let categories = category.split("\n")
        categories.forEach(element => {
          if (tagNames.includes(element)) {
            match = true;
          }
        });
      }
      if (match == true) {
        return true
      }
      match = false;
      return tagNames.includes(category);
    });
    // Return the filtered posts
    return filteredPosts;
  } else {
    // If the tag names array is empty, return the original posts
    return posts;
  }
}

// Define a function to sort the posts by the selected order
function sortPostsByOrder(posts) {
  // Get the selected order from the sort by select element
  let selectedOrder = sortBySelect.value;
  // Copy the posts array
  let sortedPosts = [...posts];
  // Sort the copied array based on the selected order
  switch (selectedOrder) {
    case "date":
      // Sort by the id in descending order (assuming the id represents the date)
      sortedPosts.sort((a, b) => b.dataset.date - a.dataset.date);
      break;
    case "likes":
      // Sort by the likes in descending order
      sortedPosts.sort((a, b) => b.dataset.likes - a.dataset.likes);
      break;
    case "dislikes":
      // Sort by the dislikes in ascending order
      sortedPosts.sort((a, b) => b.dataset.dislikes - a.dataset.dislikes);
      break;
  }
  // Return the sorted posts
  return sortedPosts;
}

// Define a function to apply the filters and the sorting to the posts
function applyFiltersAndSorting() {
  // Filter the posts by the selected tags
  let filteredPosts = filterPostsByTags(posts);
  // Sort the posts by the selected order
  let sortedPosts = sortPostsByOrder(filteredPosts);
  // Render the posts on the page
  renderFilteredPosts(sortedPosts);
}

// Add an event listener to the filter form to call the applyFiltersAndSorting function
filterForm.addEventListener("submit", event => {
  // Prevent the default form submission behavior
  event.preventDefault();
  // Call the applyFiltersAndSorting function
  applyFiltersAndSorting();
});

// Call the renderPosts function with the initial posts array
fetchPosts();

function hasSessionCookie() {
  const cookies = document.cookie.split('; ');
  return cookies.some(cookie => cookie.startsWith("sessionID"));
}