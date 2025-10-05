import targetData from './targets.json';

type TargetData = {
    [sector: string]: {
        [camera: string]: {
            [ccd: string]: any[];
        };
    };
};

interface RightPaneProps {
    sector: number
    camera: number
    ccd: number
}

const RightPane: React.FC<RightPaneProps> = (props) => {
    const currentTargets =
        (targetData as TargetData)?.[props.sector]?.[props.camera]?.[props.ccd] || [];

    return (
        <div style={{ fontFamily: 'sans-serif', padding: '20px' }}>
            <h2>CCD Target Viewer</h2>

            <h3>Top 10 Brightest Targets</h3>
            {currentTargets.length > 0 ? (
                <table style={{ width: '100%', borderCollapse: 'collapse' }}>
                    <thead style={{ textAlign: 'left' }}>
                        <tr>
                            <th style={{ border: '1px solid #ccc', padding: '8px' }}>TIC ID</th>
                            <th style={{ border: '1px solid #ccc', padding: '8px' }}>RA (deg)</th>
                            <th style={{ border: '1px solid #ccc', padding: '8px' }}>Dec (deg)</th>
                            <th style={{ border: '1px solid #ccc', padding: '8px' }}>Tmag</th>
                        </tr>
                    </thead>
                    <tbody>
                        {currentTargets.map((target) => (
                            <tr key={target.TICID}>
                                <td style={{ border: '1px solid #ccc', padding: '8px' }}>{target.TICID}</td>
                                <td style={{ border: '1px solid #ccc', padding: '8px' }}>{target.RA.toFixed(4)}</td>
                                <td style={{ border: '1px solid #ccc', padding: '8px' }}>{target.Dec.toFixed(4)}</td>
                                <td style={{ border: '1px solid #ccc', padding: '8px' }}>{target.Tmag.toFixed(2)}</td>
                            </tr>
                        ))}
                    </tbody>
                </table>
            ) : (
                <p>No data found for this selection.</p>
            )}
        </div>
    );
}

export default RightPane
