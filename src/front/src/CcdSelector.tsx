import React from 'react';

interface CcdSelectorProps {
    camera: number;
    setCamera: (n: number) => void
    ccd: number;
    setCcd: (n: number) => void
}

const CcdSelector: React.FC<CcdSelectorProps> = (props) => {
    const handleCcdClick = (cameraIndex: number, ccdIndex: number) => {
        props.setCamera(cameraIndex)
        props.setCcd(ccdIndex)
    };

    return (
        <div style={styles.cameraLayout}>
            {/* Create the 4 cameras (4x1 layout) */}
            {[1, 2, 3, 4].map((cameraNum) => {
                const isCameraSelected = props.camera === cameraNum;

                return (
                    <div
                        key={`cam-${cameraNum}`}
                        style={{
                            ...styles.camera,
                            ...(isCameraSelected ? styles.selectedCamera : {})
                        }}
                    >
                        <p style={styles.cameraTitle}>Camera {cameraNum}</p>

                        <div style={styles.ccdGrid}>
                            {[1, 2, 3, 4].map((ccdNum) => {
                                const isCcdSelected = isCameraSelected && props.ccd === ccdNum;

                                return (
                                    <div
                                        key={`ccd-${ccdNum}`}
                                        style={{
                                            ...styles.ccd,
                                            ...(isCcdSelected ? styles.selectedCcd : {})
                                        }}
                                        onClick={() => handleCcdClick(cameraNum, ccdNum)}
                                    >
                                        {ccdNum}
                                    </div>
                                );
                            })}
                        </div>
                    </div>
                );
            })}
        </div>
    );
};

const styles: { [key: string]: React.CSSProperties } = {
    cameraLayout: {
        display: 'flex',
        flexDirection: 'row',
        gap: '10px',
    },
    camera: {
        flex: 1,
        border: '2px solid #555882ff',
        borderRadius: '4px',
        padding: '8px',
        transition: 'border-color 0.2s ease',
    },
    selectedCamera: {
        border: '2px solid whitesmoke',
    },
    cameraTitle: {
        marginTop: 0,
        marginBottom: '10px',
        textAlign: 'center',
        fontWeight: 'bold',
        fontSize: '14px',
    },
    ccdGrid: {
        display: 'grid',
        gridTemplateColumns: 'repeat(2, 1fr)',
        gridTemplateRows: 'repeat(2, 1fr)',
        gap: '5px',
        height: '50px',
    },
    ccd: {
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        border: '1px solid #555882ff',
        borderRadius: '4px',
        backgroundColor: '#1c1c26',
        cursor: 'pointer',
        transition: 'background-color 0.2s ease, color 0.2s ease',
        fontWeight: 'bold',
    },
    selectedCcd: {
        backgroundColor: '#424259',
        color: 'white',
        border: '1px solid whitesmoke',
    },
};

export default CcdSelector;