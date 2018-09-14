import { DataTypes } from './types/data'
import { ImageURLs, APIURLs } from './constants/urls'

declare global {
    interface Window { ddConfig: any }
}

declare const process: any

/**
 * ddTips contains all the Tooltips functionality.
 */
module ddTips {

    /**
     * Config stores configuration options
     */
    let Config: DataTypes.Config = {
        Lang: 'en',
        Theme: 'default'
    }

    /**
     * StoreData describes the types allowed in the Store.
     */
    interface StoreData {
        [key: string]: DataTypes.Hero | DataTypes.Ability | DataTypes.Item | DataTypes.Unit
    }

    /**
     * Store keeps a track of all API response data
     * This ensures that we do not need to request data multiple times.
     */
    let Store: StoreData = {}

    let SupportedDomains: string[] = [
        'dota2.com',
        'www.dota2.com',
        'opendota.com',
        'www.opendota.com',
        'dotabuff.com',
        'www.dotabuff.com'
    ]

    export function Init() {

        let cssID: string = 'ddtip-css'
        if ( !document.getElementById(cssID) ) {
            let head: HTMLHeadElement  = document.getElementsByTagName('head')[0]
            let link: HTMLLinkElement  = document.createElement('link')
            link.id   = cssID
            link.rel  = 'stylesheet'
            link.type = 'text/css'
            link.media = 'all'
            if ( process.env.NODE_ENV === 'production' ) {
                console.log('ENV: Production')
                link.href = 'https://dota.peterbooker.com/assets/latest/ddtips.css' 
            } else {
                console.log('ENV: Development')
                link.href = 'dev/ddtips.css'
            }
            
            head.appendChild(link)
        }
        
        window.onload = () => {

            // Load Config
            if ( window.ddConfig != null ) {
                Config.Lang = (window.ddConfig.Lang == null) ? 'en' : window.ddConfig.Lang
                Config.Theme = (window.ddConfig.Theme == null) ? 'default' : window.ddConfig.Theme
            }

            let links = document.links
                        
            for ( let i = 0; i < links.length; i++ ) {
                let a = <HTMLAnchorElement>links[i]
                let attrs: any = a.attributes
                let href: string = attrs['href'].value

                var parser = document.createElement('a')
                parser.href = href

                // Check if contains a supported domain
                let idx: number = SupportedDomains.indexOf(parser.hostname)

                // Check if required data attributes were set manually
                let manual: boolean = false
                if ( a.dataset.type != null && a.dataset.item != null ) {
                    manual = true
                }

                if ( idx !== -1 || manual ) {

                    let split: string[]
                    switch (SupportedDomains[idx]) {
                        case 'dota2.com':
                        case 'www.dota2.com':
                            split = cleanPath(parser.pathname).split('/')
                            if ( split.length > 1 && split[0] === 'hero' || split[0] === 'item' ) {
                                a.setAttribute( 'data-type', split[0] )
                                a.setAttribute( 'data-item', split[1] )
                            } else {
                                continue
                            }
                            break
                        case 'opendota.com':
                        case 'www.opendota.com':
                            split = cleanPath(parser.pathname).split('/')
                            if ( split.length > 1 && split[0] === 'heroes' ) {
                                a.setAttribute( 'data-type', 'hero' )
                                a.setAttribute( 'data-item', split[1] )
                            } else {
                                continue
                            }
                            break
                        case 'dotabuff.com':
                        case 'www.dotabuff.com':
                            let valid: boolean = false
                            split = cleanPath(parser.pathname).split('/')
                            if ( split.length > 1 && split[0] === 'heroes' ) {
                                valid = true
                                a.setAttribute( 'data-type', 'hero' )
                                a.setAttribute( 'data-item', split[1].replace('-', '_') )
                            }
                            if ( split.length > 1 && split[0] === 'items' ) {
                                valid = true
                                a.setAttribute( 'data-type', 'item' )
                                a.setAttribute( 'data-item', split[1].replace('-', '_') )
                            }
                            if ( !valid ) {
                                continue
                            }
                            break
                    }

                    // Mouse Events
                    a.addEventListener('mouseenter', mouseIn)
                    a.addEventListener('mouseleave', mouseOut)
                    // Touch Events
                    //a.addEventListener('touchstart', touchStart)

                }   
            }
        }
    }

    /**
     * touchStart() runs when a relevant link is touched, if the #ddTip element does not exist it is created.
     * The anchor URL is parsed to identify the type of Tooltip and an API request made for the data, this is
     * then used to build the tooltip content.
     * 
     * @param {TouchEvent} event 
     */
    function touchStart(event: TouchEvent) {
        let el = document.getElementById('ddTip')
    }

    /**
     * mouseIn() runs when the mouse enters a relevant link, if the #ddTip element does not exist it is created.
     * The anchor URL is parsed to identify the type of Tooltip and an API request made for the data, this is
     * then used to build the tooltip content.
     * 
     * @param {MouseEvent} event 
     */
    function mouseIn(event: MouseEvent) {
        let el = document.getElementById('ddTip')

        let link = <HTMLAnchorElement>event.target
                
        if ( !el ) {
            el = document.createElement('div')
            el.id = 'ddTip'
            document.body.appendChild(el)
        }
        
        let targetElem = <HTMLAnchorElement>event.target

        let type: string = targetElem.getAttribute( 'data-type' )
        let item: string = targetElem.getAttribute( 'data-item' )
        let lang: string = Config.Lang
        let theme: string = Config.Theme

        // If not already stored, fetch data
        if ( ! ( item in Store ) ) {
            let bodyClasses: string[] = ['container', 'theme-' + Config.Theme]
            let loadingContent: string = `
                <div class="${bodyClasses.join(' ')}">
                    <div class="loading">
                        <div class="loader">Loading...</div>
                    </div>
                </div>
            `
            el.innerHTML = loadingContent
        
            getCORS(APIURLs.Base + '' + type + '/' + item + '?lang=' + lang, function(request: any) {
                    
                let response: string = request.currentTarget.response || request.target.responseText
                let data: DataTypes.Ability | DataTypes.Hero | DataTypes.Item | DataTypes.Unit
                data = JSON.parse(response)
                Store[item] = data

                switch( type ) {
                    case 'ability': {
                        let Tip: AbilityTooltip = new AbilityTooltip(<DataTypes.Ability>data, theme)
                        el = Tip.buildTip(el)
                        break
                    }
                    case 'hero': {
                        let Tip: HeroTooltip = new HeroTooltip(<DataTypes.Hero>data, theme)
                        el = Tip.buildTip(el)
                        break
                    }
                    case 'item': {
                        let Tip: ItemTooltip = new ItemTooltip(<DataTypes.Item>data, theme)
                        el = Tip.buildTip(el)
                        break
                    }
                    case 'unit': {
                        let Tip: UnitTooltip = new UnitTooltip(<DataTypes.Unit>data, theme)
                        el = Tip.buildTip(el)
                        break
                    }
                }

                positionTip(el, link)
            })
        
        } else {

            let data = Store[item]

            switch( type ) {
                case 'ability': {
                    let Tip: AbilityTooltip = new AbilityTooltip(<DataTypes.Ability>data, theme)
                    el = Tip.buildTip(el)
                    break
                }
                case 'hero': {
                    let Tip: HeroTooltip = new HeroTooltip(<DataTypes.Hero>data, theme)
                    el = Tip.buildTip(el)
                    break
                }
                case 'item': {
                    let Tip: ItemTooltip = new ItemTooltip(<DataTypes.Item>data, theme)
                    el = Tip.buildTip(el)
                    break
                }
                case 'unit': {
                    let Tip: UnitTooltip = new UnitTooltip(<DataTypes.Unit>data, theme)
                    el = Tip.buildTip(el)
                    break
                }
            }
        }

        el.className = 'active'
        positionTip(el, link)
    }
    
    /**
     * mouseOut() runs when the mouse leaves a relevant link.
     * If the #ddTip element exists, it is emptied and made invisible.
     * 
     * @param {Event} event 
     */
    function mouseOut(event: MouseEvent) {
        let el = document.getElementById('ddTip')

        if ( isMouseOverTooltip(event, el) ) {

            el.addEventListener('mouseleave', mouseOut)

        } else {

            el.removeEventListener('mouseleave', mouseOut)

            if ( el ) {
                el.className = ''
                el.innerHTML = ''
                el.innerText = ''
            }

        }
    }

    /**
     * Tooltip contains the shared structure of all Tooltips.
     * TODO: Store reference to #ddTip element in all Tooltip classes.
     */
    class Tooltip {
        element: HTMLElement
        constructor() {
        }
    }

    /**
     * HeroTooltip 
     */
    class HeroTooltip extends Tooltip {

        // Contains all Hero data
        data: DataTypes.Hero

        // Theme Name
        theme: string

        /**
         * constructor()
         * 
         * @param {DataTypes.Hero} heroData 
         * @param {string} theme
         */
        constructor( heroData: DataTypes.Hero, theme: string ) {
            super()
            this.data = heroData
            this.theme = theme
        }

        /**
         * getClass() returns a CSS class string
         * based on internal class data
         *
         * @return {String}
         */
        getClass() {
            return 'hero-' + this.data.name
        }

        /**
         * getIconClasses() returns the CSS class string for hero utility icons
         * 
         * @param {String} stat
         * @return {String}
         */
        getIconClasses(stat: string) {

            if ( this.data.attribute_primary === stat ) {
                return 'primary icon'
            } else {
                return 'icon'
            }

        }

        /**
         * buildTip() adds the relevant HTML and content to the element passed in.
         * 
         * @param {HTMLElement} el
         * @return {HTMLElement}
         */
        buildTip(el: HTMLElement) {
            let bodyClasses: string[] = ['container', 'theme-' + this.theme, 'hero', this.getClass()]
            let abilities: string = ''
            for ( let ability of this.data.abilities ) {
                abilities += `<div class="abi ability-${ability.name}"><img src="${ability.img_url}" title="${ability.title}" /><div class="name">${ability.title}</div><div class="desc">${ability.desc}</div></div>`
            }

            let tipContent: string = `
            <div class="arw"></div>
            <div class="${bodyClasses.join(' ')}">
                <div class="hdr">
                    <span class="portrait">
                        <img class="img" src="${this.data.img_url}" alt="${this.data.title} Portrait - Dota 2 Hero" title="${this.data.title}" />
                    </span>
                    <h2 class="title">${this.data.title}</h2>
                    <div class="role">
                        <span class="atk">${this.data.attack_type}</span>
                        <span class="roles">${this.data.roles.join(' - ')}</span>
                    </div>
                    <div class="stats">
                        <span class="stat-grp">
                            <img class="${this.getIconClasses('intellect')}" src="${ImageURLs.intIcon}" alt="Intelligence Icon" />
                            <span class="stat">${this.data.attribute_base_int}</span>
                        </span>
                        <span class="stat-grp">
                            <img class="${this.getIconClasses('agility')}" src="${ImageURLs.agiIcon}" alt="Agility Icon" />
                            <span class="stat">${this.data.attribute_base_agi}</span>
                        </span>
                        <span class="stat-grp">
                            <img class="${this.getIconClasses('strength')}" src="${ImageURLs.strIcon}" alt="Strength Icon" />
                            <span class="stat">${this.data.attribute_base_str}</span>
                        </span>
                    </div>
                    <div class="stats">
                        <span class="stat-grp">
                            <img class="icon atk" src="${ImageURLs.attackIcon}" alt="Attack Icon" />
                            <span class="stat">${this.data.attack_dmg_min + '-' + this.data.attack_dmg_max}</span>
                        </span>
                        <span class="stat-grp">
                            <img class="icon ms" src="${ImageURLs.speedIcon}" alt="Move Speed Icon" />
                            <span class="stat">${this.data.movement_speed}</span>
                        </span>
                        <span class="stat-grp">
                            <img class="icon arm" src="${ImageURLs.armorIcon}" alt="Armor Icon" />
                            <span class="stat">${this.data.armor}</span>
                        </span>
                    </div>
                    <div class="abilities">${abilities}</div>
                </div>
                <div class="ftr">
                    <div class="bio">${this.data.bio}</div>
                </div>
            </div>
            `

            el.innerHTML = tipContent

            return el

        }

    }

    /**
     * AbilityTooltip 
     */
    class AbilityTooltip extends Tooltip {

        // Contains all Ability data
        data: DataTypes.Ability

        // Theme Name
        theme: string

        /**
         * constructor()
         * 
         * @param {DataTypes.Ability} abilityData 
         * @param {string} theme
         */
        constructor( abilityData: DataTypes.Ability, theme: string ) {
            super()
            this.data = abilityData
            this.theme = theme
        }
        
        getClass() {
            return 'ability-' + this.data.name
        }

        /**
         * buildTip() adds the relevant HTML and content to the element passed in.
         * 
         * @param {HTMLElement} el
         * @return {HTMLElement}
         */
        buildTip(el: HTMLElement) {
            let bodyClasses: string[] = ['container', 'theme-' + this.theme, 'ability', this.getClass()]
            let cd, mc, usage: string
            if ( this.data.cd != null ) {
                cd = `<span class="cooldown"><img class="dd-icon" src="${ImageURLs.cooldownIcon}" title="Cooldown Icon" />${this.data.cd.join(' / ')}</span>`
            } else {
                cd = ''
            }
            if ( this.data.mc != null ) {
                mc = `<span class="mana"><img class="dd-icon" src="${ImageURLs.manaIcon}" title="Mana Cost Icon" />${this.data.mc.join(' / ')}</span>`
            } else {
                cd = ''
            }
            if ( mc || cd ) {
                usage = `<div class="usage">${cd}${mc}</div>`
            } else {
                usage = ``
            }

            let damage: string = ''
            if ( this.data.dmg != null ) {
                damage = `<div class="damage">${this.data.dmg}</div>`
            }

            let tipContent: string = `
            <div class="arw"></div>
            <div class="${bodyClasses.join(' ')}">
                <div class="hdr">
                    <span class="portrait">
                        <img class="img" src="${this.data.img_url}" alt="${this.data.title} - Dota 2 Ability" title="${this.data.title}" />
                    </span>
                    <h2 class="title">${this.data.title}</h2>
                    <div class="affects">${this.data.affects}</div>
                    <div class="desc">${this.data.desc}</div>
                    ${damage}
                    <div class="attributes">${this.data.attributes}</div>
                    ${usage}
                </div>
                <div class="ftr">
                    <div class="bio">${this.data.lore}</div>
                </div>
            </div>
            `
            el.innerHTML = tipContent

            return el
        }
        
    }

    /**
     * ItemTooltip 
     */
    class ItemTooltip extends Tooltip {

        // Contains all Item data
        data: DataTypes.Item

        // Theme Name
        theme: string

        /**
         * constructor()
         * 
         * @param {DataTypes.Item} itemData 
         * @param {string} theme
         */
        constructor( itemData: DataTypes.Item, theme: string ) {
            super()
            this.data = itemData
            this.theme = theme
        }
            
        getClass() {
            return 'item-' + this.data.name
        }

        /**
         * buildTip() adds the relevant HTML and content to the element passed in.
         * 
         * @param {HTMLElement} el
         * @return {HTMLElement}
         */
        buildTip(el: HTMLElement) {
            let bodyClasses: string[] = ['container', 'theme-' + this.theme, 'item', this.getClass()]
            let cd, mc, usage: string
            if ( this.data.cd != null ) {
                cd = `<span class="cooldown"><img class="dd-icon" src="${ImageURLs.cooldownIcon}" title="Cooldown Icon" />${this.data.cd.join(' / ')}</span>`
            } else {
                cd = ''
            }
            if ( this.data.mc != null ) {
                mc = `<span class="mana"><img class="dd-icon" src="${ImageURLs.manaIcon}" title="Mana Cost Icon" />${this.data.mc.join(' / ')}</span>`
            } else {
                mc = ''
            }
            if ( mc || cd ) {
                usage = `<div class="usage">${cd}${mc}</div>`
            } else {
                usage = ``
            }

            let components: string = ''
            if ( this.data.components != null && this.data.components instanceof Array ) {
                this.data.components.forEach( component => {
                    components += `<img class="component-icon" src="${component.img_url}" alt="${component.title} - Dota 2 Item" title="${component.title}" />`
                })
            }

            let level: string = ''
            if ( this.data.level != null) {
                level = `<span class="lvl">${this.data.level}</span>`
            }

            let notes: string = ''
            if ( this.data.level != null) {
                notes = `<div class="notes">${this.data.notes}</div>`
            }

            let tipContent: string = `
            <div class="arw"></div>
            <div class="${bodyClasses.join(' ')}">
                <div class="hdr">
                    <span class="portrait">
                        <img class="img" src="${this.data.img_url}" alt="${this.data.title} - Dota 2 Item" title="${this.data.title}" />
                    </span>
                    <h2 class="title">${this.data.title}${level}</h2>
                    <div class="meta"><img class="dd-icon" src="${ImageURLs.goldIcon}" title="Gold Icon" />${this.data.cost} ${components}</div>
                    <div class="desc">${this.data.desc}</div>
                    ${notes}
                    ${usage}
                </div>
                <div class="ftr">
                    <div class="bio">${this.data.lore}</div>
                </div>
            </div>
            `
            el.innerHTML = tipContent

            return el
        }
        
    }

    /**
     * UnitTooltip 
     */
    class UnitTooltip extends Tooltip {

        // Contains all Unit data
        data: DataTypes.Unit

        // Theme Name
        theme: string

        /**
         * constructor()
         * 
         * @param {DataTypes.Unit} unitData 
         * @param {string} theme
         */
        constructor( unitData: DataTypes.Unit, theme: string ) {
            super()
            this.data = unitData
            this.theme = theme
        }
            
        getClass() {
            return 'unit-' + this.data.name
        }

        /**
         * buildTip() adds the relevant HTML and content to the element passed in.
         * 
         * @param {HTMLElement} el
         * @return {HTMLElement}
         */
        buildTip(el: HTMLElement) {
            let bodyClasses: string[] = ['container', 'theme-' + this.theme, 'unit', this.getClass()]
            let tipContent: string = `
            <div class="arw"></div>
            <div class="${bodyClasses.join(' ')}">
                <div class="hdr">
                    <span class="portrait">
                        <img class="img" src="${this.data.img_url}" alt="${this.data.title} - Dota 2 Ability" />
                    </span>
                    <h2 class="title">${this.data.title}</h2>
                </div>
            </div>
            `
            el.innerHTML = tipContent

            return el
        }
        
    }

    /**
     * isMouseOverTooltip() checks if the mouse is within the bounds of the element.
     * 
     * @param {MouseEvent} event 
     * @param {HTMLElement} el 
     * @return {bool}
     */
    function isMouseOverTooltip(event: MouseEvent, el: HTMLElement) {

        let posX: number = event.clientX
        let posY: number = event.clientY

        let rect: ClientRect = el.getBoundingClientRect()

        let tipWidth: number = rect.width
        let tipHeight: number = rect.height
        let tipOffsetLeft: number = rect.left
        let tipOffsetTop: number = rect.top

        let tipMinX: number = tipOffsetLeft
        let tipMaxX: number = tipOffsetLeft + tipWidth

        let tipMinY: number = tipOffsetTop
        let tipMaxY: number = tipOffsetTop + tipHeight

        if ( posX < tipMinX || posX > tipMaxX ) {
            return false
        }

        if ( posY < tipMinY || posY > tipMaxY ) {
            return false
        }

        return true

    }

    /**
     * positionTip() positions the Tooltip element above or below the Anchor element
     * 
     * @param {HTMLElement} el 
     * @param {HTMLAnchorElement} an 
     */
    function positionTip(el: HTMLElement, an: HTMLAnchorElement) {

        let rect: ClientRect = an.getBoundingClientRect()

        let linkLeftOffset: number = rect.left
        let linkTopOffset: number = rect.top

        let linkWidth: number = rect.width
        let linkHeight: number = rect.height

        let tipWidth: number = el.offsetWidth
        //let tipHeight: number = el.offsetHeight

        let left: number = linkLeftOffset - ( tipWidth / 2 ) + ( linkWidth / 2 ) + window.scrollX
        let top: number = linkTopOffset + linkHeight + window.scrollY

        el.style.left = left + 'px'
        el.style.top = top + 'px'

    }

    // HTTP GET Request
    function getCORS(url: any, success: any) {
        var xhr = new XMLHttpRequest()
        xhr.open('GET', encodeURI(url))
        xhr.onload = success
        xhr.send()

        return xhr
    }

    /**
     * Cleans the path by removing slashes from the start and end
     * @param {string} path
     */
    function cleanPath(path: string) {
        if ( path.substr( 0, 1 ) === '/' ) {
            path = path.substr( 1, path.length )
        }

        if ( path.substr( path.length - 1, 1 ) === '/' ) {
            path = path.substr( 0, path.length - 1 )
        }

        return path
    }

}

ddTips.Init()