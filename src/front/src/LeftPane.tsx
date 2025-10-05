import { useEffect, useState } from "react"
import { TransformComponent, TransformWrapper, useControls } from "react-zoom-pan-pinch"


const markers = [
    { id: 1, x: 25, y: 40, text: 'This is Marker 1' },
    { id: 2, x: 60, y: 75, text: 'Information for Marker 2' }
];

function LeftPane() {
    const maxSector = 97
    const [sector, setSector] = useState<number>(1)

    const [transformState, setTransformState] = useState({
        scale: 1,
        positionX: 0,
        positionY: 0,
    });

    const [imageDimensions, setImageDimensions] = useState({ width: 0, height: 0 });

    useEffect(() => {
        setImageDimensions({ width: 0, height: 0 });
    }, [sector]);

    const handleMarkerClick = (marker: { text: string }) => {
        alert(`You clicked on: ${marker.text}`);
        // Here you can implement more complex logic,
        // like opening a modal or a sidebar with more details.
    };

    return (
        <div style={{ textAlign: "center" }}>
            <div style={{ display: "flex", justifyContent: "center", alignItems: "center" }}>
                {sector > 1 && <button style={{ marginRight: "20px" }} onClick={() => setSector(sector - 1)}>&lt;</button>}
                <h1>Sector {sector}</h1>
                {sector < maxSector && <button style={{ marginLeft: "20px" }} onClick={() => setSector(sector + 1)}>&gt;</button>}
            </div>
            <div style={{ width: "75%", margin: "0 auto", position: "relative", overflow: "hidden" }}>
                <TransformWrapper
                    wheel={{ smoothStep: .02 }}
                    onTransformed={(_, state) => { setTransformState(state) }}
                    key={sector}
                >
                    <Controls sector={sector} />
                    <TransformComponent wrapperStyle={{ position: 'relative', width: '100%', height: '100%' }}>
                        <img
                            key={sector}
                            src={`http://localhost:8081/downloadCCD?sector=${sector}`}
                            style={{ maxWidth: "100%" }}
                            onLoad={(ev) => {
                                const { width, height } = ev.currentTarget;
                                setImageDimensions({ width: width, height: height });
                            }}
                        />
                    </TransformComponent>
                </TransformWrapper>

                {imageDimensions.width > 0 && markers.map(marker => {
                    const initialX = (marker.x / 100) * imageDimensions.width;
                    const initialY = (marker.y / 100) * imageDimensions.height;

                    const transformedX = (initialX * transformState.scale) + transformState.positionX;
                    const transformedY = (initialY * transformState.scale) + transformState.positionY;

                    return (
                        <div
                            key={marker.id}
                            className="marker non-scaling-marker"
                            style={{
                                position: 'absolute',
                                top: `${transformedY}px`,
                                left: `${transformedX}px`,
                            }}
                            onClick={() => handleMarkerClick(marker)}
                        >
                            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="red" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
                                <line x1="18" y1="6" x2="6" y2="18"></line>
                                <line x1="6" y1="6" x2="18" y2="18"></line>
                            </svg>
                        </div>
                    );
                })}
            </div>
        </div>
    )
}

const Controls = (props: { sector: number }) => {
    const { resetTransform } = useControls();
    useEffect(() => {
        resetTransform()
    }, [props.sector])

    return (
        <></>
    );
};

export default LeftPane
