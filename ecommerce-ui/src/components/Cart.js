import React from 'react'

export default class Cart extends React.Component {
    clickHandler = () => {
        this.props.modalRef.current.style.display = "block"
    }
 
    render() {
        return (
            <div ref={this.props.cartRef} className="ui inverted black labeled icon menu">
                <button style={{"borderColor": "white", "padding": "15% 2% 2%", "width": "120%"}} className="ui button item" onClick={this.clickHandler} >
                    <h3>Cart <i className="shopping cart icon"></i></h3>
                    <div className="floating ui label blue">
                        ${this.props.total}
                    </div>
                </button>
            </div>
        )
    }
}