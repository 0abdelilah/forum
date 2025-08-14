
// Appeler au chargement
window.addEventListener('DOMContentLoaded', () => {
    if (window.location.pathname.startsWith('/auth')) {
        checkHash();
    } else if (window.location.pathname == '/') {
    }
});

// Appeler à chaque changement de hash
window.addEventListener('hashchange', checkHash);

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

function checkHash() {
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

function comment(PostUUID) {
    const content = document.getElementById("commentArea-"+PostUUID).value;

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
                showToast(responseData.error);
            }
        })
}

function loadPosts() {
    const postsContainer = document.getElementsByClassName('posts')[0];

    fetch('/load-posts')
        .then(res => res.json())
        .then(responseData => {
            if (responseData.success === "true") {
                const fragment = document.createDocumentFragment();

                responseData.posts.forEach(post => {
                    const CreationDate = post.CreationDate.substring(0, 10);

                    const postDiv = document.createElement('div');
                    postDiv.className = 'post';

                    // Build comments HTML
                    let commentsHTML = '';
                    if (post.Comments && post.Comments.length > 0) {
                        post.Comments.forEach(c => {
                            const commentDate = c.Date.substring(0, 10);
                            commentsHTML += `
                            <div class="comment">
                                <p class="comment-meta"><strong>${c.Username}</strong> • ${commentDate}</p>
                                <p class="comment-content">${c.Content}</p>
                            </div>
                            `;
                        });
                    }

                    // Build post HTML
                    postDiv.innerHTML = `
                    <div class="post-header">
                        <img src="https://icon-library.com/images/facebook-user-icon/facebook-user-icon-4.jpg" />
                        <div class="author-info">
                            <h3>${post.Username}</h3>
                            <p class="meta">${CreationDate} • ${post.Category}</p>
                        </div>
                    </div>

                    <div class="content">
                        <h3 class="post-title">${post.Title}</h3>
                        <pre>${post.Content}</pre>
                    </div>

                    <h4>Comments</h4>
                    <div class="comments">
                        ${commentsHTML}
                    </div>

                    <textarea id="commentArea-${post.PostUUID}" placeholder="Write a comment..."></textarea>
                    <button onclick="comment('${post.PostUUID}');">Post Comment</button>
                `;

                    fragment.appendChild(postDiv);
                });



                postsContainer.appendChild(fragment);
            } else {
                console.log("error", responseData)
                showToast(responseData.error, "error");
            }
        });
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