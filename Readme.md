# Pathfinder Spell Grabber

This tool was first intended to crawl the pathfinder SRD for certain spells and turn them into simple HTML spell-cards.
These cards should then be printed in A5 format to serve as quick reference during play.

It was later refined to use the spell DB (a google spreadsheet also found on the SRD website), which was more reliable
and now works 100% offline.

## How to use

* Create a file named `spells.txt` in the same folder as the executable.
* Add all the spells names you want to see turned into "spell cards" to the file. One spell per line. Empty lines are
  ignored.
* Run the program. Make sure the `spell_full.csv` is in the same folder with the executable and `spells.txt`.
* You should now have an `output` folder. It contains one `.html`file with ALL your spells and one `.html` file for each
  individual file.
* Let's assume you want to print these "cards":
    * Open your browser
    * Type in the address bar: `file:\\` followed by the file path to your `allSpells.html` file
        * If you don't know how to get the file path of the file, search for that online...
    * Adjust the format and margins as you need. Authors preference: A5, landscape, 2cm margin on the right (for
      punching holes)
    * Each spell will have a page-break after it