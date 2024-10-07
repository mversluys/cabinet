# Cabinet

This is a Wails application which is for launching arcade games via mame.

When it starts, it reads cabinet.json which contains the following

```json
{
	"mame": "<directory where mame can be found>",
	"romset": [ <roms to include> ]
}
```

The romset is then shuffled and the VideoSnaps of those roms are displayed to the user.

Pressing Button 1 (Left Control) will launch the game that's being displayed at the moment.

Pressing Right/Down will advance to the next game alphabetically and pressing Left/Up will go to the previous.

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.
