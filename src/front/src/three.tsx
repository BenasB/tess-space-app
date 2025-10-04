import * as THREE from 'three';
import { updateCameraLookAt } from "./controls"

const views = {
    main: {
        left: 0,
        bottom: 0,
        width: 1.0,
        height: 1.0,
        background: new THREE.Color().setRGB(0.5, 0.5, 0.7, THREE.SRGBColorSpace),
        camera: new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 0.1, 10000),
    },
    side: {
        left: 0,
        bottom: 0,
        width: 0.4,
        height: 0.4,
        background: new THREE.Color().setRGB(0.7, 0.5, 0.5, THREE.SRGBColorSpace),
        camera: new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 0.1, 10000),
    },
};

views.side.camera.layers.enable(1);

const scene = new THREE.Scene();
const renderer = new THREE.WebGLRenderer();
renderer.setSize(window.innerWidth, window.innerHeight)
document.body.appendChild(renderer.domElement);

views.side.camera.position.set(5, 2, 1)
views.side.camera.lookAt(views.main.camera.position)

const cameraHelper = new THREE.CameraHelper(views.main.camera)
cameraHelper.layers.set(1)
scene.add(cameraHelper);

const radius = 3
const geometry = new THREE.SphereGeometry(radius);
const material = new THREE.MeshBasicMaterial({ color: 0x00ff00, wireframe: true });
const sphere = new THREE.Mesh(geometry, material);
scene.add(sphere);

function animate() {
    updateCameraLookAt(views.main.camera);

    [views.main, views.side].forEach(view => {
        const camera = view.camera;

        const left = Math.floor(window.innerWidth * view.left);
        const bottom = Math.floor(window.innerHeight * view.bottom);
        const width = Math.floor(window.innerWidth * view.width);
        const height = Math.floor(window.innerHeight * view.height);

        renderer.setViewport(left, bottom, width, height);
        renderer.setScissor(left, bottom, width, height);
        renderer.setScissorTest(true);
        renderer.setClearColor(view.background);

        camera.aspect = width / height;
        camera.updateProjectionMatrix();

        renderer.render(scene, camera);
    });
}

renderer.setAnimationLoop(animate);