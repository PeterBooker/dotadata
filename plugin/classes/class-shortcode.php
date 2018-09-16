<?php
/**
 * Primary class file for the Shortcode.
 *
 * @package dotadata
 */

namespace DotaData;

/**
 * Class ShortCode
 */
class ShortCode {

	/**
	 * Shortcode constructor.
	 *
	 * @uses Shortcode::init()
	 *
	 * @return void
	 */
	public function __construct() {
		$this->init();
	}

	/**
	 * Shortcode ID
	 *
	 * @var integer
	 */
	private static $id = 1;

	/**
	 * Shortcode initiation.
	 *
	 * A helper function, called by `Shortcode::__construct()` to initiate actions, hooks and other features needed.
	 *
	 * @uses add_shortcode()
	 *
	 * @return void
	 */
	public function init() {
		add_shortcode( 'ddsheet', array( $this, 'process' ) );
	}

	/**
	 * Shortcode constructor.
	 *
	 * @uses DotaData::init()
	 *
	 * @param array  $atts An array containing shortcode attributes.
	 * @param string $content The content inside the shortcode.
	 *
	 * @return string
	 */
	public function process( $atts = [], $content = null ) {

		$atts = array_change_key_case( (array) $atts, CASE_LOWER );

		$atts = shortcode_atts( [
			'lang'     => 'en',
			'theme'    => 'default',
			'showlore' => true,
			'type'     => 'hero',
			'item'     => 'riki',
		], $atts, $tag );

		// do stuff.
		return '<div class="dd-datasheet">TEST DATA</div>';
	}

}
