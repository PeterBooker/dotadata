<?php
/**
 * Initialise the plugin
 *
 * @package prompress
 */

namespace DotaData;

spl_autoload_register( function ( $class ) {

	// Namespace Prefix.
	$prefix = 'DotaData\\';

	// Base Directory.
	$base_dir = DOTADATA_PATH . 'classes' . DIRECTORY_SEPARATOR;

	// Check for Prefix.
	$len = strlen( $prefix );
	if ( strncmp( $prefix, $class, $len ) !== 0 ) {
		// no, move to the next registered autoloader.
		return;
	}

	// Relative Class Name.
	$relative_class = substr( $class, $len );

	// Replace the namespace prefix with the base directory, replace namespace
	// separators with directory separators in the relative class name, append
	// with .php.
	$file = $base_dir . str_replace( '\\', DIRECTORY_SEPARATOR, 'class-' . $relative_class ) . '.php';

	// If Exists, Require.
	if ( file_exists( $file ) ) {
		require $file;
	}

});

new \DotaData\Main();
