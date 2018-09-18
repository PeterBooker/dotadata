export module DataTypes {

    export interface Config {
        Lang: string
        Theme: string
    }
    
    export interface Ability {
        id: string
        name: string
        title: string
        cd: number[]
        mc: number[]
        desc: string
        affects: string
        notes: string
        dmg: string
        attributes: string
        lore: string
        img_url: string
    }
            
    export interface Hero {
        id: string
        name: string
        title: string
        roles: string[]
        armor: number
        attack_type: string
        attack_dmg_min: number
        attack_dmg_max: number
        attack_rate: number
        attack_range: number
        attribute_primary: string
        attribute_base_str: number
        attribute_str_gain: number
        attribute_base_int: number
        attribute_int_gain: number
        attribute_base_agi: number
        attribute_agi_gain: number
        movement_speed: number
        movement_turn_rate: number
        vision: number[]
        img_url: string
        bio: string
        abilities: HeroAbility[]
        talents: HeroTalent[]
    }

    export interface HeroAbility {
        name: string
        title: string
        desc: string
        img_url: string
    }

    export interface HeroTalent {
        name:  string
        title: string
    }

    export interface Item {
        id: string
        name: string
        title: string
        cd: number[]
        mc: number[]
        cost: number
        shop_tags: string[]
        quality: string
        components: ItemComponent[]
        sideshop: boolean
        secretshop: boolean
        desc: string
        notes: string
        lore: string
        level: number
        img_url: string
    }

    export interface ItemComponent {
        name: string
        title: string
        cost: number
        img_url: string
    }

    export interface Unit {
        id: string
        name: string
        title: string
        img_url: string
    }
            
}