<?php
/**
 * Primary class file for the DotaData plugin.
 *
 * @package dotadata
 */

namespace DotaData;

/**
 * Class Main
 */
class Main {

	/**
	 * DotaData constructor.
	 *
	 * @uses DotaData::init()
	 *
	 * @return void
	 */
	public function __construct() {
		$this->init();
	}

	/**
	 * Plugin initiation.
	 *
	 * A helper function, called by `DotaData::__construct()` to initiate actions, hooks and other features needed.
	 *
	 * @uses add_action()
	 * @uses add_filter()
	 *
	 * @return void
	 */
	public function init() {
		add_action( 'wp_loaded', array( $this, 'load_i18n' ) );
		add_action( 'wp_enqueue_scripts', array( $this, 'enqueue_files' ) );
		add_action( 'plugins_loaded', array( $this, 'shortcode' ) );

		// Register Block if Gutenberg is available.
		if ( function_exists( 'gutenberg_init' ) ) {
			add_action( 'plugins_loaded', array( $this, 'block' ) );
		}
	}

	/**
	 * Load translations.
	 *
	 * Loads the textdomain needed to get translations for our plugin.
	 *
	 * @uses load_plugin_textdomain()
	 * @uses basename()
	 * @uses dirname()
	 *
	 * @return void
	 */
	public function load_i18n() {
		load_plugin_textdomain( 'dotadata', false, basename( dirname( __FILE__ ) ) . '/languages/' );
	}

	/**
	 * Enqueue assets.
	 *
	 * Conditionally enqueue our CSS and JavaScript when viewing plugin related pages in wp-admin.
	 *
	 * @uses wp_enqueue_style()
	 * @uses plugins_url()
	 * @uses wp_enqueue_script()
	 * @uses wp_localize_script()
	 * @uses esc_html__()
	 *
	 * @return void
	 */
	public function enqueues() {
		wp_enqueue_style( 'dd-sheet', DOTADATA_URL . '/assets/css/ddsheet.min.css', array(), DOTADATA_VERSION );
		wp_enqueue_script( 'dd-tooltips', DOTADATA_URL . '/assets/js/ddtips.min.js', array(), DOTADATA_VERSION, true );
		wp_localize_script( 'dd-tooltips', 'ddConfig', array(
			'Lang'          => 'en', // TODO: Integrate with i18n plugins.
			'Theme'         => 'default',
			'ShowLore'      => true,
			'IncludeStyles' => true,
		) );
	}

	/**
	 * Register Shortcode.
	 *
	 * Registers the Shortcode for the plugin.
	 *
	 * @uses \DotaData\Shortcode()
	 *
	 * @return void
	 */
	public function shortcode() {
		new \DotaData\Shortcode();
	}

	/**
	 * Register Block.
	 *
	 * Registers the Block for the plugin.
	 *
	 * @uses \DotaData\Block()
	 *
	 * @return void
	 */
	public function block() {
		new \DotaData\Block();
	}

}
