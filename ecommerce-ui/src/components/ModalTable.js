import React from 'react'

import ModalTableCell from './ModalTableCell'

export default class ModalTable extends React.Component {
    render() {
        return (
            <table style={{ "width": "95%", "textAlign": "center" }} className="ui very basic collapsing celled table">
                <thead>
                    <tr>
                        <th>Item</th>
                        <th>Price</th>
                    </tr>
                </thead>
                <tbody>
                    {[...Array.from(this.props.list.keys())].map(key => {
                        const count = this.props.list.get(key)
                        const product = this.props.products[key-1]

                        return <ModalTableCell key={key} count={count} product={product}/>
                    })}

                    <tr>
                        <td><h3>Total Cost</h3></td>
                        <td>${this.props.total}</td>
                    </tr>
                </tbody>
            </table>
        )
    }
}