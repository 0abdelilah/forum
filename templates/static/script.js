////////////////////
// Authentication //
////////////////////

window.addEventListener('DOMContentLoaded', checkHash);
window.addEventListener('hashchange', checkHash)

// if hash is #login -> display login & hide register and vice versa 
function checkHash() {
    if (window.location.pathname.startsWith('/auth')) {

        hash = window.location.hash;

        const loginForm = document.getElementById('loginForm');
        const registerForm = document.getElementById('registerForm');

        if (hash === "#register") {
            registerForm.style.setProperty('display', 'block');
            loginForm.style.setProperty('display', 'none', 'important');
        } else {
            loginForm.style.setProperty('display', 'block');
            registerForm.style.setProperty('display', 'none', 'important');
        }
    }
}

function Login() {
    username = document.getElementById('username').value
    password = document.getElementById('password').value

    fetch('/login', {
        method: 'POST',
        body: JSON.stringify({ username, password }),
    })
        .then(res => res.json())
        .then(responseData => {
            if (responseData.success === "true") {
                window.location.href = '/';
            } else {
                showToast(responseData.error, "error");
            }
        })
        .catch(err => {
            console.error('Login error:', err);
        });
}

function Register() {
    username = document.getElementById('r_username').value
    email = document.getElementById('email').value
    password = document.getElementById('r_password').value

    fetch('/register', {
        method: 'POST',
        body: JSON.stringify({ username, email, password }),
    })
        .then(res => res.json())
        .then(responseData => {
            if (responseData.success === "true") {
                window.location.hash = '#login';
                showToast("Registered");
            } else {
                showToast(responseData.error, "error");
            }
        })
        .catch(err => {
            console.error('Login error:', err);
        });
}

// will display based on param set by setCategory
function formatDate(dateString) {
    return dateString.substring(0, 10);
}

function renderComment(comment) {
    return `
        <div class="comment">
            <p class="comment-meta"><strong>${comment.Username}</strong> • ${formatDate(comment.Date)}</p>
            <p class="comment-content">${comment.Content}</p>
        </div>
    `;
}

function renderPost(post) {
    const commentsHTML = (post.Comments || []).map(renderComment).join('');

    const postDiv = document.createElement('div');
    postDiv.className = 'post';
    postDiv.innerHTML = `
        <div class="post-header">
            <img src="https://icon-library.com/images/facebook-user-icon/facebook-user-icon-4.jpg" />
            <div class="author-info">
                <h3>${post.Username}</h3>
                <p class="meta">${formatDate(post.CreationDate)} • ${post.Category}</p>
            </div>
        </div>

        <div class="content">
            <h3 class="post-title">${post.Title}</h3>
            <pre>${post.Content}</pre>
        </div>

        <h4>Comments</h4>
        <div class="comments">${commentsHTML}</div>

        <textarea id="commentArea-${post.PostUUID}" placeholder="Write a comment..."></textarea>
        <button class="comment-btn" data-id="${post.PostUUID}">Post Comment</button>
        
        <div class="react">
            <button class="like-btn" data-id="${post.PostUUID}">${post.Reacts.Likes.Amount} - Like</button>
            <button class="dislike-btn" data-id="${post.PostUUID}">${post.Reacts.Dislikes.Amount} - Dislike</button>
        </div>
    `;
    return postDiv;
}

// render like & dislike buttons (unused)
function renderSvg(react) {
    if (react == 'like') {
        return
    }
}

// helper for loadPosts
// sets category param given by argument 
function setCategory(category) {
    const url = new URL(window.location);

    if (category === "all") {
        url.searchParams.delete("category"); // Remove the category param entirely
    } else {
        url.searchParams.set("category", category);
    }

    // Update the address bar without reloading
    history.replaceState(null, "", url);

    loadPosts();
}

// helper function to extact name from html
function getUsername() {
    const name = document.querySelector('label[for="show"]').dataset.name;
    if (name) {
        return name;
    }
    showToast("Login to see this category", "error")
}

function showSelected() {
    const params = new URLSearchParams(window.location.search);
    let category = params.get("category");

    // hide all
    const radios = document.querySelectorAll(`input[type="radio"]`);

    radios.forEach(radio => {

        // hide all
        const icon = radio.nextElementSibling.querySelector('svg.icon');
        icon.style.display = 'none';

        if (!category) {
            category = "all"
        }

        const checked = document.querySelector(`input[type="radio"][value="${category}"]`);
        const checkedicon = checked.nextElementSibling.querySelector('svg.icon');
        checkedicon.style.display = 'inline-block';
    })
}

function loadPosts() {

    showSelected()
    const postsContainer = document.querySelector('.posts');
    postsContainer.innerHTML = '';

    const params = new URLSearchParams(window.location.search);
    let category = params.get("category");

    if (category) {
        const allowed = ["mine", "liked", "programming", "music", "gaming"];
        if (!allowed.includes(category)) {
            showToast("Category doesn't exist", "error");
            category = null;
        }
    }

    let username;
    if (category === "mine" || category === "liked") {
        username = getUsername();
        if (!username) return;
    }

    fetch('/load-posts')
        .then(res => res.json())
        .then(data => {
            if (data.success !== "true") {
                showToast(data.error, "error");
                return;
            }

            const posts = data.posts || [];
            let filteredPosts = posts;

            if (category === "mine") {
                filteredPosts = posts.filter(p => p.Username === username);
            } else if (category === "liked") {
                // posts dont have likes (null)
                filteredPosts = posts.filter(post => (post.Reacts.Likes.Usernames || []).includes(username));
            } else if (category) {
                filteredPosts = posts.filter(p => p.Category === category);
            }

            if (!filteredPosts.length) {
                postsContainer.textContent = "No posts available.";
                return;
            }

            const fragment = document.createDocumentFragment(); // create a mini dom 
            filteredPosts.forEach(post => fragment.appendChild(renderPost(post))); // append posts to it
            postsContainer.appendChild(fragment);

            // Event listeners
            postsContainer.querySelectorAll('.comment-btn').forEach(btn =>
                btn.addEventListener('click', () => comment(btn.dataset.id))
            );
            postsContainer.querySelectorAll('.like-btn').forEach(btn =>
                btn.addEventListener('click', () => react(btn.dataset.id, 'like'))
            );
            postsContainer.querySelectorAll('.dislike-btn').forEach(btn =>
                btn.addEventListener('click', () => react(btn.dataset.id, 'dislike'))
            );
        })
        .catch(err => {
            console.error("Load posts error:", err);
            showToast("Could not load posts", "error");
        });
}

function react(PostUUID, reaction) {
    fetch("/react", {
        method: 'POST',
        body: JSON.stringify({ PostUUID, reaction }),
    })
        .then(res => res.json())
        .then(responseData => {
            if (responseData.success == "true") {
                showToast("Reacted to post");
                loadPosts();
            } else {
                console.log(responseData.error)
                showToast(responseData.error, "error");
            }
        })
}

function CreatePost() {
    title = document.getElementById('title').value
    content = document.getElementById('content').value
    category = document.getElementById('category').value

    fetch("/create-post", {
        method: "POST",
        body: JSON.stringify({ title, content, category }),
    })
        .then(res => res.json())
        .then(responseData => {
            if (responseData.success === "true") {
                window.location.reload();
                showToast("Created post");
            } else {
                showToast(responseData.error, "error");
            }
        })
        .catch(err => {
            console.error('Login error:', err);
        });
}

function comment(PostUUID) {
    const content = document.getElementById("commentArea-" + PostUUID).value;

    fetch('/comment', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ content, PostUUID })
    })
        .then(res => res.json())
        .then(responseData => {
            if (responseData.success === "true") {
                window.location.reload();
                showToast("Commented successfully");
            } else {
                showToast(responseData.error, "error");
            }
        })
}

function showToast(message, state = "success") {
    // state can be "success" or "error"
    const toast = document.createElement('div');
    toast.id = state === "success" ? "ToastSuccess" : "ToastError";
    toast.textContent = message;

    // append to dashboard or authmenu depending on what's visible
    const dashboardPage = document.getElementsByClassName('dashboard-page')[0];
    const authMenu = document.getElementsByClassName('authmenu')[0];
    if (dashboardPage) {
        dashboardPage.appendChild(toast);
    } else if (authMenu) {
        authMenu.appendChild(toast);
    }
}