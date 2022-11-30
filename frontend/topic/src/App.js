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
    const data = [
        {category: "Sporting Goods", price: "$49.99", stocked: true, name: "Football"},
        {category: "Sporting Goods", price: "$9.99", stocked: true, name: "Baseball"},
        {category: "Sporting Goods", price: "$29.99", stocked: false, name: "Basketball"},
        {category: "Electronics", price: "$99.99", stocked: true, name: "iPod Touch"},
        {category: "Electronics", price: "$399.99", stocked: false, name: "iPhone 5"},
        {category: "Electronics", price: "$199.99", stocked: true, name: "Nexus 7"}
    ]
    return (
        <div className="App">
            <FilterableProductTable data={data}/>
        </div>
    );
}

export default App;
const styles = {
    font1: {
        color: '#FF0000',
    },

    font2: {
        color: '#2D9900',
    },
}

class ProductRow extends React.Component {
    render() {
        const row = this.props.row
        console.log("row " + row)
        return (
            <div>
                {row.stocked
                    ? <h3>{row.name}</h3>
                    : <h3 style={styles.font1}>{row.name}</h3>
                }
                <h4>{row.price.toString()}</h4>
            </div>
        )
    }
}

class ProductCategoryRow extends React.Component {
    render() {
        const rows = this.props.rows
        return (
            <div>
                {rows.map(row => <ProductRow row={row}/>)}
            </div>
        )
    }
}

function groupBy(xs, keyGetter) {
    return xs.reduce(function (rv, x) {
        (rv[keyGetter(x)] = rv[keyGetter(x)] || []).push(x);
        return rv;
    }, {});
}

class ProductTable extends React.Component {
    render() {
        const categorys = groupBy(this.props.data, row => row.category)

        console.log("aaaa" + categorys)
        console.log("aaaa1 " + categorys['Electronics'])

        return (
            <ul>
                {Object.keys(categorys).map(k => {
                    console.log("aaaa3" + k)
                    return (
                        <div>
                            <label>{k}</label>
                            <ProductCategoryRow rows={categorys[k]}/>
                        </div>
                    )
                })}
            </ul>
        )
    }
}

class SearchBar extends React.Component {
    render() {
        return (
            <div>
                <input type='text' onChange={e => this.props.handleSearch(e.target.value)}/>
                <input type='checkbox' onChange={e => this.props.handleOnlyStock(e.target.checked)}/>
            </div>
        )
    }
}

class FilterableProductTable extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            onlyStock: false,
            data: props.data,
            search: '',
        }
        this.handleOnlyStock = this.handleOnlyStock.bind(this)
        this.handleSearch = this.handleSearch.bind(this)
        this.filter = this.filter.bind(this)
    }

    filter() {
        const search = this.state.search.toLowerCase()
        const data = this.props.data.filter(item => {
            return (this.state.search === '' || item.name.toLowerCase().includes(search))
                && !(this.state.onlyStock && !item.stocked)
        })
        this.setState({data: data})
    }

    handleOnlyStock(checked) {
        this.setState({onlyStock: checked})
        this.filter()
    }

    handleSearch(s) {
        this.setState({search: s})
        this.filter()
    }

    render() {
        const filterData = this.state.data
        return (
            <div>
                <SearchBar handleOnlyStock={this.handleOnlyStock} handleSearch={this.handleSearch}/>
                <ProductTable data={filterData}/>
            </div>
        )
    }
}
