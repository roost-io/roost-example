import React from 'react'

import Cart from './Cart'
import Modal from './Modal'
import ProductList from './ProductList'

export default class App extends React.Component {
    modalRef = React.createRef()
    cartRef = React.createRef()

    state = {
        total: 0,
        list: new Map(),
        products: []
    }

    componentDidMount() {
        fetch(`http://${(process.env.REACT_APP_API_URL) || 'localhost:8080'}/api/products`)
            .then(response => response.json())
            .then(data => this.setState({
                products: data
            }))
            .catch(err => console.error(err))
    }

    cartHandler = (id, shouldInc, currPrice) => {
        this.setState(({total, list}) => {
            const newList = new Map(list)
            const newTotal = total + (currPrice)*(shouldInc ? 1 : -1)

            if (shouldInc) {
                newList.set(id, list.get(id) ? list.get(id) + 1 : 1)
            } else {
                if (list.get(id) === 1) {
                    newList.delete(id)
                } else {
                    newList.set(id, list.get(id) - 1)
                }
            }

            return {list: newList, total: newTotal}
        })
    }

    hideModal = (e) => {
        if(! this.cartRef.current.contains(e.target)) {
            if (! this.modalRef.current.contains(e.target)) {
                this.modalRef.current.style.display = "none"
            }
        }
    }

    render() {
        return (
            <div className="ui inverted black segment fluid container" onClick={this.hideModal}>
                <h1 style={{"fontSize": "300%"}} className="ui blue header center aligned">TechMart</h1>
                <div className="ui top attached tabular menu">
                    <div className="active item">Catalogue</div>
                    <Cart cartRef={this.cartRef} total={this.state.total} modalRef={this.modalRef} />
                    <Modal total={this.state.total} products={this.state.products} list={this.state.list} modalRef={this.modalRef} />
                </div>
                <div className="ui inverted black bottom attached active tab segment">
                    <ProductList products={this.state.products} cartHandler={this.cartHandler} />
                </div>
            </div>
        )
    }
}