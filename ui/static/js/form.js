// Neumorphism Login Form JavaScript
class NeumorphismLoginForm {
    constructor() {
        this.form = document.getElementById('loginForm');
        this.emailInput = document.getElementById('email');
        this.passwordInput = document.getElementById('password');
        this.passwordToggle = document.getElementById('passwordToggle');
        this.submitButton = this.form.querySelector('.login-btn');
        this.successMessage = document.getElementById('successMessage');
        this.socialButtons = document.querySelectorAll('.neu-social');
        
        this.init();
    }
    
    init() {
        this.bindEvents();
        this.setupPasswordToggle();
        this.setupSocialButtons();
        this.setupNeumorphicEffects();
    }
    
    bindEvents() {
        this.form.addEventListener('submit', (e) => this.handleSubmit(e));
        this.emailInput.addEventListener('blur', () => this.validateEmail());
        this.passwordInput.addEventListener('blur', () => this.validatePassword());
        this.emailInput.addEventListener('input', () => this.clearError('email'));
        this.passwordInput.addEventListener('input', () => this.clearError('password'));
        
        // Add soft press effects to inputs
        [this.emailInput, this.passwordInput].forEach(input => {
            input.addEventListener('focus', (e) => this.addSoftPress(e));
            input.addEventListener('blur', (e) => this.removeSoftPress(e));
        });
    }
    
    setupPasswordToggle() {
        this.passwordToggle.addEventListener('click', () => {
            const type = this.passwordInput.type === 'password' ? 'text' : 'password';
            this.passwordInput.type = type;
            
            this.passwordToggle.classList.toggle('show-password', type === 'text');
            
            // Add soft click animation
            this.animateSoftPress(this.passwordToggle);
        });
    }
    
    setupSocialButtons() {
        this.socialButtons.forEach(button => {
            button.addEventListener('click', (e) => {
                this.animateSoftPress(button);
                
                // Determine which social platform based on SVG content
                const svgPath = button.querySelector('svg path').getAttribute('d');
                let provider = 'Social';
                if (svgPath.includes('22.56')) provider = 'Google';
                else if (svgPath.includes('github')) provider = 'GitHub';
                else if (svgPath.includes('23.953')) provider = 'Twitter';
                
                this.handleSocialLogin(provider, button);
            });
        });
    }
    
    setupNeumorphicEffects() {
        // Add hover effects to all neumorphic elements
        const neuElements = document.querySelectorAll('.neu-icon, .neu-checkbox, .neu-social');
        neuElements.forEach(element => {
            element.addEventListener('mouseenter', () => {
                element.style.transform = 'scale(1.05)';
            });
            
            element.addEventListener('mouseleave', () => {
                element.style.transform = 'scale(1)';
            });
        });
        
        // Add ambient light effect on mouse move
        document.addEventListener('mousemove', (e) => {
            this.updateAmbientLight(e);
        });
    }
    
    updateAmbientLight(e) {
        const card = document.querySelector('.login-card');
        const rect = card.getBoundingClientRect();
        const x = e.clientX - rect.left;
        const y = e.clientY - rect.top;
        
        const centerX = rect.width / 2;
        const centerY = rect.height / 2;
        
        const angleX = (x - centerX) / centerX;
        const angleY = (y - centerY) / centerY;
        
        const shadowX = angleX * 30;
        const shadowY = angleY * 30;
        
        card.style.boxShadow = `
            ${shadowX}px ${shadowY}px 60px #bec3cf,
            ${-shadowX}px ${-shadowY}px 60px #ffffff
        `;
    }
    
    addSoftPress(e) {
        const inputGroup = e.target.closest('.neu-input');
        inputGroup.style.transform = 'scale(0.98)';
    }
    
    removeSoftPress(e) {
        const inputGroup = e.target.closest('.neu-input');
        inputGroup.style.transform = 'scale(1)';
    }
    
    animateSoftPress(element) {
        element.style.transform = 'scale(0.95)';
        setTimeout(() => {
            element.style.transform = 'scale(1)';
        }, 150);
    }
    
    validateEmail() {
        const email = this.emailInput.value.trim();
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        
        if (!email) {
            this.showError('email', 'Email is required');
            return false;
        }
        
        if (!emailRegex.test(email)) {
            this.showError('email', 'Please enter a valid email');
            return false;
        }
        
        this.clearError('email');
        return true;
    }
    
    validatePassword() {
        const password = this.passwordInput.value;
        
        if (!password) {
            this.showError('password', 'Password is required');
            return false;
        }
        
        if (password.length < 6) {
            this.showError('password', 'Password must be at least 6 characters');
            return false;
        }
        
        this.clearError('password');
        return true;
    }
    
    showError(field, message) {
        const formGroup = document.getElementById(field).closest('.form-group');
        const errorElement = document.getElementById(`${field}Error`);
        
        formGroup.classList.add('error');
        errorElement.textContent = message;
        errorElement.classList.add('show');
        
        // Add gentle shake animation
        const input = document.getElementById(field);
        input.style.animation = 'gentleShake 0.5s ease-in-out';
        setTimeout(() => {
            input.style.animation = '';
        }, 500);
    }
    
    clearError(field) {
        const formGroup = document.getElementById(field).closest('.form-group');
        const errorElement = document.getElementById(`${field}Error`);
        
        formGroup.classList.remove('error');
        errorElement.classList.remove('show');
        setTimeout(() => {
            errorElement.textContent = '';
        }, 300);
    }
    
    async handleSubmit(e) {
    e.preventDefault();
    
    const isEmailValid = this.validateEmail();
    const isPasswordValid = this.validatePassword();
    
    if (!isEmailValid || !isPasswordValid) return;

    this.setLoading(true);

    try {
        const response = await fetch('http://localhost:8080/api/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                email: this.emailInput.value,
                password: this.passwordInput.value
            })
        });

        const data = await response.json();

        if (response.ok) {
            // Successful login from Go server
            this.showNeumorphicSuccess();
        } else {
            // Server returned an error (e.g., 401 Unauthorized)
            this.showError('password', data.message || 'Invalid credentials');
        }
    } catch (error) {
        // Network error or server is down
        this.showError('password', 'Server connection failed');
    } finally {
        this.setLoading(false);
    }
}
    
    async handleSocialLogin(provider, button) {
        console.log(`Initiating ${provider} login...`);
        
        // Add loading state to button
        button.style.pointerEvents = 'none';
        button.style.opacity = '0.7';
        
        try {
            await new Promise(resolve => setTimeout(resolve, 1500));
            console.log(`Redirecting to ${provider} authentication...`);
            // window.location.href = `/auth/${provider.toLowerCase()}`;
        } catch (error) {
            console.error(`${provider} authentication failed: ${error.message}`);
        } finally {
            button.style.pointerEvents = 'auto';
            button.style.opacity = '1';
        }
    }
    
    setLoading(loading) {
        this.submitButton.classList.toggle('loading', loading);
        this.submitButton.disabled = loading;
        
        // Disable social buttons during login
        this.socialButtons.forEach(button => {
            button.style.pointerEvents = loading ? 'none' : 'auto';
            button.style.opacity = loading ? '0.6' : '1';
        });
    }
    
    showNeumorphicSuccess() {
        // Soft fade out form
        this.form.style.transform = 'scale(0.95)';
        this.form.style.opacity = '0';
        
        setTimeout(() => {
            this.form.style.display = 'none';
            document.querySelector('.social-login').style.display = 'none';
            document.querySelector('.signup-link').style.display = 'none';
            
            // Show success with soft animation
            this.successMessage.classList.add('show');
            
            // Animate success icon
            const successIcon = this.successMessage.querySelector('.neu-icon');
            successIcon.style.animation = 'successPulse 0.6s ease-out';
            
        }, 300);
        
        // Simulate redirect
        setTimeout(() => {
            console.log('Redirecting to dashboard...');
            // window.location.href = '/dashboard';
        }, 2500);
    }
}

class NeumorphismSignupForm extends NeumorphismLoginForm {
    constructor() {
        super();
        this.form = document.getElementById('signupForm');
        this.usernameInput = document.getElementById('username');
        this.confirmPasswordInput = document.getElementById('confirmPassword');
    }

    init() {
        super.init();
        this.usernameInput.addEventListener('blur', () => this.validateUsername());
        this.usernameInput.addEventListener('input', () => this.clearError('username'));
        this.confirmPasswordInput.addEventListener('blur', () => this.validateConfirmPassword());
        this.confirmPasswordInput.addEventListener('input', () => this.clearError('confirmPassword'));

        [this.usernameInput, this.confirmPasswordInput].forEach(input => {
            input.addEventListener('focus', (e) => this.addSoftPress(e));
            input.addEventListener('blur', (e) => this.removeSoftPress(e));
        });
    }

    validateUsername() {
        const username = this.usernameInput.value.trim();
        if (!username) {
            this.showError('username', 'Username is required');
            return false;
        }
        if (username.length < 3) {
            this.showError('username', 'Username must be at least 3 characters');
            return false;
        }
        this.clearError('username');
        return true;
    }

    validateConfirmPassword() {
        const password = this.passwordInput.value;
        const confirmPassword = this.confirmPasswordInput.value;

        if (!confirmPassword) {
            this.showError('confirmPassword', 'Please confirm your password');
            return false;
        }
        if (password !== confirmPassword) {
            this.showError('confirmPassword', 'Passwords do not match');
            return false;
        }
        this.clearError('confirmPassword');
        return true;
    }

    async handleSubmit(e) {
        e.preventDefault();
        
        const isUsernameValid = this.validateUsername();
        const isEmailValid = this.validateEmail();
        const isPasswordValid = this.validatePassword();
        const isConfirmValid = this.validateConfirmPassword();
        
        if (!isUsernameValid || !isEmailValid || !isPasswordValid || !isConfirmValid) return;

        this.setLoading(true);

        try {
            const response = await fetch('http://localhost:8080/api/signup', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    username: this.usernameInput.value,
                    email: this.emailInput.value,
                    password: this.passwordInput.value
                })
            });

            const data = await response.json();

            if (response.ok) {
                this.showNeumorphicSuccess();
            } else {
                this.showError('email', data.message || 'Registration failed');
            }
        } catch (error) {
            this.showError('password', 'Server connection failed');
        } finally {
            this.setLoading(false);
        }
    }

    showNeumorphicSuccess() {
        this.form.style.transform = 'scale(0.95)';
        this.form.style.opacity = '0';
        
        setTimeout(() => {
            this.form.style.display = 'none';
            const social = document.querySelector('.social-login');
            const link = document.querySelector('.signup-link');
            const divider = document.querySelector('.divider');
            if(social) social.style.display = 'none';
            if(link) link.style.display = 'none';
            if(divider) divider.style.display = 'none';
            
            this.successMessage.classList.add('show');
            const successIcon = this.successMessage.querySelector('.neu-icon');
            successIcon.style.animation = 'successPulse 0.6s ease-out';
        }, 300);
        
        setTimeout(() => {
            window.location.href = 'login.html';
        }, 2500);
    }
}

// Add custom animations
if (!document.querySelector('#neu-keyframes')) {
    const style = document.createElement('style');
    style.id = 'neu-keyframes';
    style.textContent = `
        @keyframes gentleShake {
            0%, 100% { transform: translateX(0); }
            25% { transform: translateX(-3px); }
            75% { transform: translateX(3px); }
        }
        
        @keyframes successPulse {
            0% { transform: scale(0.8); opacity: 0; }
            50% { transform: scale(1.1); }
            100% { transform: scale(1); opacity: 1; }
        }
    `;
    document.head.appendChild(style);
}

// Initialize the form when DOM is loaded
document.addEventListener('DOMContentLoaded', () => {
    if (document.getElementById('loginForm')) {
        new NeumorphismLoginForm();
    } else if (document.getElementById('signupForm')) {
        new NeumorphismSignupForm();
    }
});