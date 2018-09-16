<?php
/**
 * Plugin Name: DotaData
 * Plugin URI: https://github.com/PeterBooker/dotadata/
 * Description: Easily display Dota2 data on your website.
 * Version: 0.1.0
 * Author: Peter Booker
 * Author URI: https://www.peterbooker.com
 *
 * @package dotadata
 */

// Useful global constants.
define( 'DOTADATA_VERSION', '0.7.0' );
define( 'DOTADATA_URL', plugin_dir_url( __FILE__ ) );
define( 'DOTADATA_PATH', plugin_dir_path( __FILE__ ) );
define( 'DOTADATA_FILE', plugin_basename( __FILE__ ) );

dd_pre_init();

/**
 * Verify that we can initialize the PromPress, then load it.
 *
 * @since 0.1.0
 */
function dd_pre_init() {

	// Check WP Version
	// Get unmodified $wp_version.
	include ABSPATH . WPINC . '/version.php';
	// Strip '-src' from the version string. Messes up version_compare().
	$version = str_replace( '-src', '', $wp_version );
	if ( version_compare( $version, '4.8', '<' ) ) {
		add_action( 'admin_notices', 'dotadata_wordpress_version_notice' );
		return;
	}

	/**
	 * Load plugin initialisation file.
	 */
	require plugin_dir_path( __FILE__ ) . '/init.php';

}

/**
 * Display a WP version notice and deactivate the PromPress plugin.
 *
 * @since 0.1.0
 */
function dotadata_wordpress_version_notice() {
	echo '<div class="error"><p>';
	echo 'DotaData requires WordPress 4.8 or later to function properly. Please upgrade WordPress before activating DotaData.';
	echo '</p></div>';
	deactivate_plugins( array( 'dotadata/dotadata.php' ) );
}
