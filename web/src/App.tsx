import { BrowserRouter as Router, Route } from 'react-router-dom';
import { ReactComponent as Logo } from './assets/images/logo.svg';
import { Characters } from './app/rickandmorty-characters/characters';
import './App.css';
import "./index.css"

function App() {
  return (
    <div className="App">
      <Router>
        <header className="App-header">
          <Logo fill="white"></Logo>
          <main>
            <nav>
              <ul>
                <li><a href="/">Characters</a></li>
                <li><a href="/Links">Environment links</a></li>
              </ul>
            </nav>
          </main>
        </header>
        <div className="App-content">
          <Route path="/" exact  component={Characters} />
        </div>
      </Router>

    </div>
  );
}

export default App;

