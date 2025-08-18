## **Header / Navigation Bar** //

* **Logout button** (only visible when logged in)
* **Login & Register buttons** (only visible when logged out)
* **Filter menu / dropdown**

  * By category (available to everyone)
  * By "My Posts" (only for logged-in users)
  * By "Liked Posts" (only for logged-in users)

---

## **Home / Main Page** //

* **List of posts** (visible to everyone)

  * Post title, author, date, category tags
  * Like & dislike counters
  * Like / dislike buttons (only for logged-in users)

---

## **Post Page** (single post view)

* Full post content
* Post categories
* Like/dislike system for the post (only for logged-in users)
* **Comments section**

  * List of comments (visible to everyone)
  * Like/dislike system for comments (only for logged-in users)
  * Add comment form (only for logged-in users)

---

## **Authentication Pages**   //

* **Login page**

  * Email
  * Password
  * Error display for wrong credentials
* **Register page**

  * Email (error if already taken)
  * Username 
  * Password (encrypted in DB)
  * Password strength check

---

## **Post Creation Page** (only logged-in users)

* Title
* Content
* Category selection (multi-select possible)
* Submit button

---

## **Backend Technical Features**

* SQLite database
* Sessions & cookies for authentication
* Password encryption with bcrypt
* UUID for session IDs
* Docker containerization
* HTTP error handling & validation

