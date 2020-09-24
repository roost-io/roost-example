import React from 'react'

import Product from './Product'

export default class ProductList extends React.Component {
    render() {
        return (
            <div className="ui three column grid">
                {this.props.products.map(product => <Product cartHandler={this.props.cartHandler} product={product} key={product.id} />)}
            </div>
        )
    }
}