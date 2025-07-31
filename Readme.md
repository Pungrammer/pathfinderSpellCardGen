# Pathfinder Spell Card Generator

This tool was first intended to crawl the pathfinder SRD for certain spells and turn them into simple HTML spell-cards.
These cards should then be printed in A5 format to serve as quick reference during play.

It was later refined to use the [spell DB](https://www.d20pfsrd.com/magic/tools/spells-db/), which was more reliable
and now works 100% offline.

## How to use

* Use the commands `list fields` and `list options <fieldname>` to craft a filter. For details see the `Filter` section.
* Run the command `export <filter>`.
* You should now have an `output` directory. It contains one `.html`file with ALL your spells and one `.html` file for each spell.
* Let's assume you want to print these "cards":
    * Open your browser.
    * Type in the address bar: `file:\\` followed by the file path to your `allSpells.html` file
        * If you don't know how to get the file path of the file, search for that online...
    * Adjust the format and margins as you need. Authors preference: A5, landscape, 2 cm margin on the left (for
      punching holes).
    * Each spell will have a page-break after it.

## Filter

The filter is an [expr-lang](https://expr-lang.org/docs/language-definition).
The expression must evaluate to `true` or `false`.

## Commands

### list


To see which fields are available, use `./pathfinderSpellCardGen list fields`.
To see a list of options for fields (e.g., which sources exist) use `./pathfinderSpellCardGen list options <fieldname>`

### export

Example filter to get all spell cards for a level 6 paladin using a subset of sources:  
`./pathfinderSpellCardGen export 'Name == "Detect Evil" || (PaladinLevel == "1" && (Source contains "Ultimate" || Source in ["PFRPG Core", "APG", "Advanced Class Guide"]))'`  
I had to add "Detect Evil" explicitly, since it is normally not on the spell list, but paladins get the ability to cast it as a class feature.

### filter

If you want to see which spells match a filter, you can use the `filter` command, which prints a list of spell names, which matched the given filter.  
`./pathfinderSpellCardGen filter 'SorcererLevel == "0"'`  
This filter would give you all sourcerer cantrips, regardless of source.

### print

To further refine what you want to see, you can use the `print` command.  
This command outputs a subset of `fields` which you define.  
`./pathfinderSpellCardGen print 'Name' 'SorcererLevel == "0"'`  
This command is similar to `filter`, but the format is a bit different.
`Name` is a bit redundant, but the print command needs at least one field.
If you want to print everything, you can supply the special value `'ALL'`:  
`./pathfinderSpellCardGen print 'ALL' 'SorcererLevel == "0"'`  



## Planned improvements

* More advanced parsing to assign data types correctly:
  * Some columns have "0" and "1" to indicate boolean values. These should become actual booleans (e.g. `CostlyComponents`).
  * Some columns are numeric only, but use "NULL" as a null indicator. These should also be parsed into a number type (e.g. `PaladinLevel`).
* Build a web server based on this tool
