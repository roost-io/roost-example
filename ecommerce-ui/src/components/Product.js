import React from 'react'

import AddBtn from './buttons/AddBtn'
import IncDcrBtn from './buttons/IncDcrBtn'

export default class Product extends React.Component{
    state = {
        count: 0,
        imageUrl: null
    }

    componentDidMount() {
        fetch(`http://${(process.env.REACT_APP_API_URL) || 'localhost:8080'}/api/products/image/${this.props.product.imageUrl}`)
            .then(data => data.blob())
            .then(image => this.setState({imageUrl: URL.createObjectURL(image)}))
            .catch(err => console.error(err))
    }

    componentWillUnmount() {
        if (this.state.imageUrl) {
            URL.revokeObjectURL(this.state.imageUrl)
        }
    }

    countUpdateHandler = (shouldInc) => {
        this.setState(prevState => ({
            count: prevState.count + (shouldInc ? 1 : -1)
        }))
        
        const {id, price} = this.props.product
        this.props.cartHandler(id, shouldInc, price)
    }

    render() {
        const product = this.props.product
        
        return (
            <div className="column" key={product.id}>
                <div style={{"width": "95%"}} key={product.id} className="ui card">
                    <div className="image">
                        <img src={this.state.imageUrl} alt={product.name} />
                    </div>
                    {this.state.count !== 0 ? <IncDcrBtn count={this.state.count} countHandler={this.countUpdateHandler} /> : <AddBtn countHandler={this.countUpdateHandler}/>}
                    <div className="floating ui label blue">
                        ${product.price}
                    </div>
                </div>
            </div>
        )
    }
}