// main.js

document.addEventListener('DOMContentLoaded', function() {
    // Select all artist tiles within the artist grid
    const artistTiles = document.querySelectorAll('.artist-grid .artist-card');

    // Add click event listener to each artist tile
    artistTiles.forEach(tile => {
        tile.addEventListener('click', function(event) {
            // Prevent the default action if the click was not on a link
            if (!event.target.closest('a')) {
                event.preventDefault();
                
                // Get the artist ID from the tile's data attribute
                const artistId = this.dataset.artistId;
                
                // Construct the URL for the artist detail page
                const artistUrl = `/artist/${artistId}`;
                
                // Navigate to the artist detail page
                window.location.href = artistUrl;
            }
        });

        // Add hover effect
        tile.addEventListener('mouseenter', function() {
            this.style.transform = 'scale(1.05)';
            this.style.transition = 'transform 0.3s ease-in-out';
        });

        tile.addEventListener('mouseleave', function() {
            this.style.transform = 'scale(1)';
        });
    });
});