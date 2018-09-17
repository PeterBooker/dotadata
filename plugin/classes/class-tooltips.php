<?php
/**
 * Class file for the Tooltips.
 *
 * @package dotadata
 */

namespace DotaData;

/**
 * Class ShortCode
 */
class Tooltips {

	/**
	 * Block constructor.
	 *
	 * @uses Tooltips::init()
	 *
	 * @return void
	 */
	public function __construct() {
		$this->init();
	}

	/**
	 * Tooltips initiation.
	 *
	 * @uses add_shortcode()
	 *
	 * @return void
	 */
	public function init() {
		add_action( 'wp_enqueue_scripts', array( $this, 'enqueue' ) );
	}

	/**
	 * Block Editor Assets.
	 *
	 * Enqueues Gutenberg Block editor assets.
	 *
	 * @return void
	 */
	public function enqueue() {
		if ( is_page() || is_singular( 'post' ) ) {
			wp_enqueue_script( 'dd-tooltips', 'https://dota.peterbooker.com/assets/latest/ddtips.js', array(), DOTADATA_VERSION, true );
			wp_localize_script( 'dd-tooltips', 'ddConfig', array(
				'Lang'          => 'no', // TODO: Integrate with i18n plugins.
				'Theme'         => 'default',
				'ShowLore'      => true,
				'IncludeStyles' => true,
			) );
		}
	}

}
