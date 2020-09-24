import React from 'react'

export default class AddBtn extends React.Component {
    btnRef = React.createRef()

    clickHandler = () => {
        this.props.countHandler(true)
        this.btnRef.current.style.display="None"
    }

    render() {
        return (
            <button className="fluid ui blue right labeled icon button" onClick={this.clickHandler} ref={this.btnRef}>
                Add to cart
                <i className="shopping cart icon white"></i>
            </button>
        )
    }
}