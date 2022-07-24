import React from 'react';
import content from './data/data.json';

function App() {
  return (
    <div>
      {JSON.stringify(content)}
    </div>
  );
}

export default App;
