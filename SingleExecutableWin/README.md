# How to create single executable file in Windows

* Open run propmt (Win+R)
* run `iexpress`
* follow the wizard installation:
  * Create new Self Exraction Directive file -> Next
  * Extract files and run an installation command -> Next
  * <add package name> -> Next
  * <select prompt or not> -> Next
  * <select show license or not> -> Next
  * Add -> <select files that are part of the single executable> -> Next
  * Add Install Program script or command -> Next
	* **Note** to make such executable, that runs some script, then under Install Program write: cmd /c <script name>. **NB!** first create dummy package, to see what the name of the script is inside the package (iexpress might not keep the original name). For example, my script was named `CONTRO~1.BAT`.
  * Default(recommended) -> Next
  * <select finish message or not> -> Next
  * <select package path and name> -> tick additional wanted Options (eg Hide File Extracting Process Animation from User) -> Next
  * <select restart configuration> -> Next
  * <Save Self Extraction file or not> **Note** not necessary -> Next
  * Create Package -> Next
  * Finish

## Author

Written by
Meelis Utt