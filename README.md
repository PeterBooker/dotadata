# DotaData

**note: The project is still being finalised and the public website is mostly still placeholder data. It should all be finished by this weekend.**

A group of projects which help you display translated Dota2 data on your website, beautifully. Add stunning tooltips to dota2.com, opendota.com, dotabuff.com and custom links. Or embed Dota2 data (Heroes, Abilities and Items) directly into webpages with the brilliant WordPress plugin. If that is not enough, use the API directly and create your own content.

## Demo

You can find more information and view a demo of the tooltips at [https://dota.peterbooker.com/](https://dota.peterbooker.com/).

## Tooltips

The tooltips script lives in the `/tooltips/` folder. It contains a TypeScript project which creates `.js` and `.css` files which can be added to your website, automatically making relevant links display tooltips. You can customize the output by writing a config to the page, e.g. `<script>var ddConfig = {Lang: 'en', Theme: 'default', ShowLore: false};</script>`.

## WordPress Plugin

The WordPress plugin lives in the `/plugin/` folder. It is under development and should be finished this week.

## API

The API collects information into convenient groups for Abilities, Heroes and Items. It is publicly available at `https://dota.peterbooker.com/api/v1/`.

### Languages

The API endpoints which return data for a single entity (hero, ability, item) all accept accept the `lang` URL parameter with the short language code. Only works with languages supported by Dota2. Example request: `https://dota.peterbooker.com/api/v1/hero/zeus?lang=no` (Hero Zeus, Language Norwegian).

### Heroes

List of all available Heroes:
`https://dota.peterbooker.com/api/v1/heroes`

Specific Hero:
`https://dota.peterbooker.com/api/v1/hero/{name|id}`

### Abilities

List of all available Abilities:
`https://dota.peterbooker.com/api/v1/abilities`

Specific Ability:
`https://dota.peterbooker.com/api/v1/ability/{name}`

### Items

List of all available Items:
`https://dota.peterbooker.com/api/v1/items`

Specific Item:
`https://dota.peterbooker.com/api/v1/item/{name}`