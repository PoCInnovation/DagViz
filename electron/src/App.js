import './App.css';
import content from './data/output.json';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <p>
          {JSON.stringify(content)}
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;
