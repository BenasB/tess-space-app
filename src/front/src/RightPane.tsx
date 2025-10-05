

interface RightPaneProps {
    sector: number
    camera: number
    ccd: number
    targets: any[]
}

const RightPane: React.FC<RightPaneProps> = (props) => {

    return (
        <div style={{ fontFamily: 'sans-serif', padding: '20px' }}>
            <h3>Top 10 Brightest Targets</h3>
            {props.targets.length > 0 ? (
                <table style={{ width: '100%', borderCollapse: 'collapse', border: "none" }}>
                    <thead style={{ textAlign: 'left' }}>
                        <tr>
                            <th style={{ border: '1px solid #ccc', padding: '8px' }}>TIC ID</th>
                            <th style={{ border: '1px solid #ccc', padding: '8px' }}>RA (deg)</th>
                            <th style={{ border: '1px solid #ccc', padding: '8px' }}>Dec (deg)</th>
                            <th style={{ border: '1px solid #ccc', padding: '8px' }}>Tmag</th>
                        </tr>
                    </thead>
                    <tbody>
                        {props.targets.map((target) => (
                            <tr key={target.TICID}>
                                <td style={{ border: '1px solid #ccc', padding: '8px' }}>
                                    <a
                                        target="_blank" rel="noopener noreferrer"
                                        href={`https://mast.stsci.edu/portal/Mashup/Clients/Mast/Portal.html?searchQuery=%7B%22service%22%3A%22TIC%22%2C%22inputText%22%3A%22TIC%20${target.TICID}%22%2C%22paramsService%22%3A%22Mast.Catalogs.Tic.Cone%22%2C%22title%22%3A%22TICv8.2%3A%20TIC%20${target.TICID}%22%2C%22columns%22%3A%22*%22%7D`}>
                                        {target.TICID}
                                    </a>
                                </td>
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
