import { useEffect, useMemo, useState } from "react"
import { TransformComponent, TransformWrapper, useControls } from "react-zoom-pan-pinch"
import CcdSelector from "./CcdSelector";

interface LeftPaneProps {
    sector: number
    setSector: (sector: number) => void
    camera: number
    setCamera: (camera: number) => void
    ccd: number
    setCcd: (ccd: number) => void
    targets: any[]
}

const LeftPane: React.FC<LeftPaneProps> = ({ sector, camera, ccd, setSector, setCamera, setCcd, targets }) => {
    const maxSector = 97

    const [transformState, setTransformState] = useState({
        scale: 1,
        positionX: 0,
        positionY: 0,
    });

    const [imageDimensions, setImageDimensions] = useState({ width: 0, height: 0 });

    const markers: { id: number, x: number, y: number, text: string }[] = useMemo(() => {
        if (!targets[0].x_percent) {
            return []
        }
        return targets.map(t => ({ id: t.TICID, x: 100 * t.pixel_x / 2136, y: 100 * t.pixel_y / 2078, text: "lol" }))
    }, [targets, imageDimensions])

    useEffect(() => {
        setImageDimensions({ width: 0, height: 0 });
    }, [sector]);

    const handleMarkerClick = (marker: { text: string }) => {
        alert(`You clicked on: ${marker.text}`);
    };

    return (
        <div style={{ textAlign: "center" }}>
            <div style={{ display: "flex", justifyContent: "center", alignItems: "center" }}>
                {sector > 1 && <button style={{ marginRight: "20px" }} onClick={() => setSector(sector - 1)}>&lt;</button>}
                <h1>Sector {sector}</h1>
                {sector < maxSector && <button style={{ marginLeft: "20px" }} onClick={() => setSector(sector + 1)}>&gt;</button>}
            </div>
            <div style={{ width: "430px", margin: "15px auto" }}>
                <CcdSelector camera={camera} setCamera={setCamera} ccd={ccd} setCcd={setCcd} />
            </div>
            <div style={{ width: "75%", margin: "0 auto", position: "relative", overflow: "hidden" }}>
                <TransformWrapper
                    wheel={{ smoothStep: .02 }}
                    onTransformed={(_, state) => { setTransformState(state) }}
                    key={`${sector}-${camera}-${ccd}`}
                >
                    <Controls sector={sector} />
                    <TransformComponent wrapperStyle={{ position: 'relative', width: '100%', height: '100%' }}>
                        <img
                            key={sector}
                            src={`http://localhost:8081/downloadCCD?sector=${sector}&camera=${camera}&ccd=${ccd}`}
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
