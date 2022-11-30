import logo from './logo.svg';
import './App.css';
import React from 'react';

class LoginBoard extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            username: "",
            password: "",
        }
    }

    render() {
        return (
            <div>
                <label>
                    <input type="text" onChange={e => this.setState({'username': e.target.value})}/>
                </label>
                <label>
                    <input type="password" onChange={e => this.setState({'password': e.target.value})}/>
                </label>
                <div>
                    <button type="submit" onClick={() => {
                        fetch("http://localhost:8080/login", {
                            method: 'POST',
                            headers: {
                                Accept: 'application/json',
                                'Content-Type': 'application/json',
                            },
                            body: JSON.stringify(this.state)
                        })
                            .then(res => res.json())
                            .then(data => {
                                    console.log(data)
                                },
                                (err) => {
                                    console.error("aa", err.message)
                                }
                            )
                    }
                    }>提交
                    </button>
                </div>
            </div>
        )
    }
}

function App() {
    return (
        <div className="App">
            <LoginBoard/>
        </div>
    );
}

export default App;
