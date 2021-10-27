import { BrowserRouter as Router, Route } from 'react-router-dom';
import { ReactComponent as Logo } from './assets/images/logo.svg';
import { Characters } from './app/rickandmorty-characters/characters';
import { Locations } from './app/rickandmorty-locations/locations';
import './App.css';
import "./index.css"

function App() {
  return (
    <div className="App">
      <Router>
        <header className="App-header">
          <main>
            <nav>
              <ul className="App-menu">
                <Logo className="App-menu-item--logo"></Logo>
                <li className="App-menu-item"><a href="/Characters">Characters</a></li>
                <li className="App-menu-item"><a href="/Locations">Locations</a></li>
              </ul>
            </nav>
          </main>
        </header>
        <div className="App-content">
          <Route path="/" exact component={Characters} />
          <Route path="/Characters" exact component={Characters} />
          <Route path="/Locations" exact component={Locations} />
        </div>
      </Router>

    </div>
  );
}

export default App;

