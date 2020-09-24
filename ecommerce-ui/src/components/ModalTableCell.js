import React from 'react'

export default class ModalTableCell extends React.Component {
    state = {
        imageUrl: null
    }

    componentDidMount() {
        fetch(`http://${(process.env.REACT_APP_API_URL) || 'localhost:8080'}/api/products/image/${this.props.product.imageUrl}`)
            .then(data => data.blob())
            .then(image => this.setState({imageUrl: URL.createObjectURL(image)}))
            .catch(err => console.error(err))
    }

    render() {
        return (
            <tr>
                <td>
                    <h4 className="ui image header">
                        <img alt={this.props.product.name} src={this.state.imageUrl} className="ui image" />
                        <div className="content">
                            {this.props.product.name}
                            <div className="sub header">${this.props.product.price}</div>
                        </div>
                    </h4>
                </td>
                <td>{this.props.count} x ${this.props.product.price}</td>
            </tr>
        )
    }
}