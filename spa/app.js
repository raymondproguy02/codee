const display = document.querySelector('#display');
        const regBtn = document.getElementById('regBtn');
        const logBtn = document.getElementById('logBtn');
        const homeBtn = document.getElementById('homeBtn');
        const aboutBtn = document.getElementById('aboutBtn');
        const contactBtn = document.getElementById('contactBtn');
        const servicesBtn = document.getElementById('servicesBtn');

        // Define content with more details
        const home = `
            <h2>🏠 Home Page</h2>
            <p>Welcome to our Single Page Application!</p>
            <p>Use the navigation buttons above to switch between pages.</p>
            <p style="color: #666; margin-top: 20px;">💡 This is a simple SPA example</p>
        `;

        const signup = `
            <h2>📝 Sign Up</h2>
            <p>Create your account</p>
            <form id="signupForm">
                <input type="text" placeholder="Full Name" required>
                <input type="email" placeholder="Email Address" required>
                <input type="password" placeholder="Password" required>
                <input type="password" placeholder="Confirm Password" required>
                <button type="submit">Register Now</button>
            </form>
        `;

        const signin = `
            <h2>🔑 Sign In</h2>
            <p>Welcome back!</p>
            <form id="signinForm">
                <input type="email" placeholder="Email Address" required>
                <input type="password" placeholder="Password" required>
                <button type="submit">Login</button>
            </form>
            <p style="margin-top: 15px; color: #666;">
                Don't have an account? <a href="#" id="switchToSignup">Sign Up</a>
            </p>
        `;

        // New pages you added
        const about = `
            <h2>ℹ️ About Us</h2>
            <p>We're building amazing Single Page Applications!</p>
            <p>Our mission: Make web development simple and fun.</p>
            <p style="margin-top: 15px; color: #666;">⭐ We're a team of passionate developers</p>
        `;

        const contact = `
            <h2>📧 Contact Us</h2>
            <p>Email: support@myapp.com</p>
            <p>Phone: +1 (555) 123-4567</p>
            <p>Follow us on social media!</p>
            <div class="social-links">
                <a href="#">📘 Facebook</a>
                <a href="#">🐦 Twitter</a>
                <a href="#">📸 Instagram</a>
                <a href="#">💼 LinkedIn</a>
            </div>
        `;

        const services = `
            <h2>🛠️ Our Services</h2>
            <ul>
                <li>🌐 Web Development</li>
                <li>📱 Mobile Apps</li>
                <li>☁️ Cloud Solutions</li>
                <li>🎨 UI/UX Design</li>
                <li>🤖 AI & Machine Learning</li>
                <li>🔒 Cybersecurity</li>
            </ul>
            <p style="margin-top: 15px; color: #666;">Contact us for custom solutions!</p>
        `;

        // Set initial content
        display.innerHTML = home;

        // Event listeners for all buttons
        homeBtn.addEventListener('click', function() {
            display.innerHTML = home;
        });

        aboutBtn.addEventListener('click', function() {
            display.innerHTML = about;
        });

        servicesBtn.addEventListener('click', function() {
            display.innerHTML = services;
        });

        contactBtn.addEventListener('click', function() {
            display.innerHTML = contact;
        });

        regBtn.addEventListener('click', function() {
            display.innerHTML = signup;
        });

        logBtn.addEventListener('click', function() {
            display.innerHTML = signin;
            
            // Add switch to signup functionality
            const switchLink = document.getElementById('switchToSignup');
            if (switchLink) {
                switchLink.addEventListener('click', function(e) {
                    e.preventDefault();
                    display.innerHTML = signup;
                });
            }
        });

        // Handle form submissions with alert
        document.addEventListener('submit', function(e) {
            e.preventDefault();
            if (e.target.id === 'signupForm') {
                alert('✅ Account created successfully! (Demo)');
                display.innerHTML = home;
            } else if (e.target.id === 'signinForm') {
                alert('✅ Logged in successfully! (Demo)');
                display.innerHTML = home;
            }
        });

        console.log('🚀 App is ready with 6 pages!');
