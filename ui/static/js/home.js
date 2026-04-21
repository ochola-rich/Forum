{/* <script type="module">
        import { initializeApp } from "https://www.gstatic.com/firebasejs/11.6.1/firebase-app.js";
        import { getAuth, signInAnonymously, onAuthStateChanged } from "https://www.gstatic.com/firebasejs/11.6.1/firebase-auth.js";
        import { getFirestore, collection, addDoc, onSnapshot, query, serverTimestamp } from "https://www.gstatic.com/firebasejs/11.6.1/firebase-firestore.js";

        // Configuration from Environment
        const firebaseConfig = JSON.parse(__firebase_config);
        const app = initializeApp(firebaseConfig);
        const auth = getAuth(app);
        const db = getFirestore(app);
        const appId = typeof __app_id !== 'undefined' ? __app_id : 'forum-demo';

        let currentUser = null;

        // Initialize Auth (Rule 3)
        async function initAuth() {
            try {
                await signInAnonymously(auth);
            } catch (err) {
                console.error("Auth failed", err);
            }
        }

        onAuthStateChanged(auth, (user) => {
            currentUser = user;
            if (user) {
                setupFeedListener();
            }
        });

        // Setup Firestore Listener (Rule 1 & 2)
        function setupFeedListener() {
            const postsRef = collection(db, 'artifacts', appId, 'public', 'data', 'posts');
            
            // Simple query only (Rule 2)
            onSnapshot(postsRef, (snapshot) => {
                const container = document.getElementById('posts-container');
                container.innerHTML = '';

                if (snapshot.empty) {
                    container.innerHTML = '<div class="loading-text">No posts yet. Be the first to start a conversation!</div>';
                    return;
                }

                // Sort in memory (Rule 2 requirement)
                const posts = [];
                snapshot.forEach(doc => posts.push({ id: doc.id, ...doc.data() }));
                posts.sort((a, b) => (b.createdAt?.seconds || 0) - (a.createdAt?.seconds || 0));

                posts.forEach(post => {
                    const card = createPostCard(post);
                    container.appendChild(card);
                });
            }, (error) => {
                console.error("Snapshot error:", error);
                document.getElementById('posts-container').innerHTML = '<div class="loading-text">Error loading feed. Please refresh.</div>';
            });
        }

        function createPostCard(post) {
            const div = document.createElement('article');
            div.className = 'post-card';
            
            const avatarColor = post.authorColor || '#3b82f6';
            const initial = (post.authorName || 'User').charAt(0).toUpperCase();

            div.innerHTML = `
                <div class="post-header">
                    <div class="user-meta">
                        <div class="avatar" style="background-color: ${avatarColor}">${initial}</div>
                        <div class="user-details">
                            <span class="name" style="display:block; font-weight:600">${post.authorName || 'Anonymous'}</span>
                            <span class="time" style="color:var(--text-muted); font-size:12px">${formatDate(post.createdAt)}</span>
                        </div>
                    </div>
                    <div class="post-card-tags">
                        <span class="tag-dark">${post.category || 'General'}</span>
                    </div>
                </div>
                <h3 class="post-title">${escapeHTML(post.title)}</h3>
                <p class="post-body">${escapeHTML(post.body)}</p>
                <div class="post-footer">
                    <div class="interaction"><span>👍</span> ${post.likes || 0}</div>
                    <div class="interaction"><span>👎</span> ${post.dislikes || 0}</div>
                    <div class="interaction"><span>💬</span> ${post.commentCount || 0} Comments</div>
                </div>
            `;
            return div;
        }

        // Global functions for UI interaction
        window.openModal = () => {
            document.getElementById('modal-overlay').style.display = 'flex';
        };

        window.closeModal = (e) => {
            if (e && e.target !== e.currentTarget) return;
            document.getElementById('modal-overlay').style.display = 'none';
        };

        window.submitPost = async () => {
            if (!currentUser) return;

            const title = document.getElementById('post-title-input').value;
            const body = document.getElementById('post-body-input').value;
            const category = document.getElementById('post-category-input').value;

            if (!title || !body) return;

            const colors = ['#14b8a6', '#3b82f6', '#c56cf0', '#f43f5e', '#f59e0b'];

            try {
                const postsRef = collection(db, 'artifacts', appId, 'public', 'data', 'posts');
                await addDoc(postsRef, {
                    title,
                    body,
                    category: category || 'General',
                    authorName: 'Guest User',
                    authorColor: colors[Math.floor(Math.random() * colors.length)],
                    likes: 0,
                    dislikes: 0,
                    commentCount: 0,
                    createdAt: serverTimestamp()
                });
                
                // Clear inputs and close
                document.getElementById('post-title-input').value = '';
                document.getElementById('post-body-input').value = '';
                document.getElementById('post-category-input').value = '';
                closeModal();
            } catch (err) {
                console.error("Error adding post:", err);
            }
        };

        function formatDate(timestamp) {
            if (!timestamp) return 'Just now';
            const date = timestamp.toDate();
            return date.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' }) + ' - ' + date.toLocaleDateString();
        }

        function escapeHTML(str) {
            const p = document.createElement('p');
            p.textContent = str;
            return p.innerHTML;
        }

        initAuth();
    </script> */}