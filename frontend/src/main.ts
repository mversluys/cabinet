import './style.css';
import './app.css';

import {GetRomset, GetMachines, Play} from '../wailsjs/go/main/App';

const startSound = new Audio('assets/start.wav');

(async () => {
    let romset = await GetRomset();
    let machines = await GetMachines();
    let sorted = [...romset].sort((a, b) => machines[a].Description.localeCompare(machines[b].Description));
    let shuffled = [...romset].sort(() => Math.random() - 0.5);
    let shuffledToSorted: { [key: string]: number } = {};
    shuffled.forEach(rom => { shuffledToSorted[rom] = sorted.indexOf(rom); });
    let current = 0;

    function update(rom: string) {
        let machine = machines[rom];
        (document.getElementById("title") as HTMLSpanElement).textContent = machine.Description;
        (document.getElementById("year") as HTMLSpanElement).textContent = machine.Year;
        (document.getElementById("manufacturer") as HTMLSpanElement).textContent = machine.Manufacturer;
        (document.getElementById("players") as HTMLSpanElement).textContent = machine.Input.Players;
        resetAnimation(document.getElementById("information") as HTMLElement);

        const videoPlayer = document.getElementById("video-player") as HTMLVideoElement;
        videoPlayer.src = `http://localhost:8080/video?name=${rom}`;
        videoPlayer.onended = next;
    }

    function resetAnimation(element: HTMLElement) {
        element.style.animation = 'none';
        element.offsetHeight;
        element.style.animation = '';
    }

    function next() {
        current = (current + 1) % shuffled.length;
        update(shuffled[current]);
    }

    function navigate(next: boolean) {
        const i = shuffledToSorted[shuffled[current]];
        let sortedIndex;
        if (next) {
            sortedIndex = (i + 1) % sorted.length;
        } else {
            sortedIndex = (i - 1 + sorted.length) % sorted.length;
        }
        const rom = sorted[sortedIndex];
        current = shuffled.indexOf(rom); 
        update(rom);
    }

    update(shuffled[current]);

    document.addEventListener('keydown', async (event) => {
        if (event.key === "ArrowUp" || event.key === "ArrowLeft") {
            event.preventDefault();
            navigate(false);
        } else if (event.key === "ArrowDown" || event.key === "ArrowRight") {
            event.preventDefault();
            navigate(true);
        } else if (event.code === "ControlLeft") {

            startSound.play();

            const videoPlayer = document.getElementById("video-player") as HTMLVideoElement;
            videoPlayer.pause();
            videoPlayer.style.opacity = "0";

            await Play(shuffled[current]);

            videoPlayer.style.opacity = "1";
            videoPlayer.play();

            await new Promise(resolve => setTimeout(resolve, 1000));
        }
    });
})();

declare global {
    interface Window {
    }
}
