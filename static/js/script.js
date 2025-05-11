document.addEventListener('DOMContentLoaded', function() {
    // Set current year in footer - no longer needed since we're using server-side year
    
    // Initialize Prism.js for syntax highlighting
    if (typeof Prism !== 'undefined') {
        Prism.highlightAll();
    }
    
    // Add smooth scrolling for anchor links
    document.querySelectorAll('a[href^="#"]').forEach(anchor => {
        anchor.addEventListener('click', function(e) {
            e.preventDefault();
            
            const targetId = this.getAttribute('href');
            if (targetId === '#') return;
            
            const target = document.querySelector(targetId);
            if (target) {
                window.scrollTo({
                    top: target.offsetTop - 80, // Adjust for header height
                    behavior: 'smooth'
                });
            }
        });
    });
    
    // Add copy button to code blocks
    document.querySelectorAll('pre code').forEach(block => {
        // Create copy button
        const button = document.createElement('button');
        button.className = 'copy-button';
        button.textContent = 'Copy';
        
        // Position the button
        const pre = block.parentNode;
        pre.style.position = 'relative';
        button.style.position = 'absolute';
        button.style.top = '0.5rem';
        button.style.right = '0.5rem';
        button.style.padding = '0.25rem 0.5rem';
        button.style.fontSize = '0.8rem';
        button.style.background = '#4a4a4a';
        button.style.color = 'white';
        button.style.border = 'none';
        button.style.borderRadius = '3px';
        button.style.cursor = 'pointer';
        
        // Add copy functionality
        button.addEventListener('click', () => {
            const code = block.textContent;
            navigator.clipboard.writeText(code).then(() => {
                button.textContent = 'Copied!';
                setTimeout(() => {
                    button.textContent = 'Copy';
                }, 2000);
            }).catch(err => {
                console.error('Failed to copy code: ', err);
                button.textContent = 'Error!';
                setTimeout(() => {
                    button.textContent = 'Copy';
                }, 2000);
            });
        });
        
        pre.appendChild(button);
    });
    
    // Mobile navigation toggle
    const createMobileNav = () => {
        if (window.innerWidth <= 768) {
            const nav = document.querySelector('nav');
            if (nav && !document.querySelector('.mobile-nav-toggle')) {
                const toggle = document.createElement('button');
                toggle.className = 'mobile-nav-toggle';
                toggle.innerHTML = '&#9776;'; // hamburger icon
                toggle.style.background = 'none';
                toggle.style.border = 'none';
                toggle.style.fontSize = '1.5rem';
                toggle.style.cursor = 'pointer';
                toggle.style.display = 'none'; // Hide initially until we add functionality
                
                nav.parentNode.insertBefore(toggle, nav);
                
                // We're not implementing the full mobile nav in this example
                // In a real app, you'd toggle the nav visibility here
            }
        }
    };
    
    // Call once on load
    createMobileNav();
    
    // Check on resize
    window.addEventListener('resize', createMobileNav);
});
