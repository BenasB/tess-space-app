import { useEffect, useState } from 'react'

function App() {
  const [count, setCount] = useState(0)
  const [helloMessage, setHelloMessage] = useState<string | null>(null)

  useEffect(() => {
    async function startFetching() {
      setHelloMessage(null);
      const response = await fetch("http://localhost:8081");
      if (ignore || !response.ok) {
        return
      }
      const body = await response.bytes()
      setHelloMessage(new TextDecoder().decode(body));
    }

    let ignore = false;
    startFetching();
    return () => {
      ignore = true;
    }
  }, [count]);

  return (
    <>
      <h1>tess-space-app</h1>
      <div>
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
      </div>
      <div>
        {helloMessage && <>
          <p>API response: {helloMessage}</p>
        </>}
      </div>
    </>
  )
}

export default App
