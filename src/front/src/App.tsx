import "./App.css"
import LeftPane from './LeftPane';
import RightPane from './RightPane';

function App() {
  return (
    <div className="split-screen-container">
      <div className="left-pane">
        <LeftPane></LeftPane>
      </div>
      <div className="right-pane">
        <RightPane></RightPane>
      </div>
    </div>
  )
}

export default App
