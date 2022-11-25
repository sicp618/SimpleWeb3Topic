import logo from './logo.svg';
import './App.css';

global.User = {
    username: "aaron",
    password: "111111",
}

function App() {
    return (
        <div className="App">
            <header className="App-header">
                <label>
                    <input type="text" onChange={e => global.User.username = e.target.value}/>
                </label>
                <label>
                    <input type="password" onChange={e => this.state.password = e.target.value}/>
                </label>
                <div>
                    <button type="submit" onClick={e => {
                        fetch("http://localhost:8080/login", {
                            method: 'POST',
                            body: JSON.stringify(global.User),
                            mode: 'no-cors',
                            headers: {
                                'Content-type': 'application/json; charset=UTF=8',
                                'Accept': '*/*',
                            },
                        })
                            .then(res => res.json())
                            .then(data => {
                                console.log(data)
                            })
                            .catch(err => {
                                console.error(err.message)
                            })
                    }}>提交
                    </button>
                </div>
            </header>
        </div>
    );
}

export default App;
