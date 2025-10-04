import * as THREE from 'three';

let isDragging = false;
let previousMousePosition = { x: 0, y: 0 };
let phi = 0;
let theta = Math.PI / 2;
const speed = 0.005

const raTextElement = document.getElementById("ra")
const decTextElement = document.getElementById("dec")

document.addEventListener('mousedown', () => {
    isDragging = true;
});

function myMod(a: number, n: number) {
    return ((a % n) + n) % n
}

document.addEventListener('mousemove', (event) => {
    if (isDragging) {
        let deltaX = event.clientX - previousMousePosition.x;
        let deltaY = event.clientY - previousMousePosition.y;

        phi -= deltaX * speed
        phi = myMod(phi, 2 * Math.PI);
        theta -= deltaY * speed;

        // Limit vertical angle (theta) to avoid flipping the camera upside down
        theta = Math.max(Math.min(theta, Math.PI - 0.01), 0.01);

        const dec = 90 - (theta * 180 / Math.PI);
        let ra = phi * 180 / Math.PI;
        const raHours = ra / 15;

        if (raTextElement) {
            raTextElement.textContent = raHours.toFixed(4)
        }
        if (decTextElement) {
            decTextElement.textContent = dec.toFixed(4)
        }
    }

    previousMousePosition = { x: event.clientX, y: event.clientY };
});

document.addEventListener('mouseup', () => {
    isDragging = false;
});

export function updateCameraLookAt(camera: THREE.Camera) {
    const targetX = Math.sin(theta) * Math.cos(phi);
    const targetY = Math.cos(theta);
    const targetZ = Math.sin(theta) * Math.sin(phi);

    camera.lookAt(targetX, targetY, targetZ);
}