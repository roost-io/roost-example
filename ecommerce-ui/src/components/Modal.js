import React from 'react'

import ModalTable from './ModalTable'

export default class Modal extends React.Component {
    render() {
        return (
            <div ref={this.props.modalRef} style={{ "width": "75%", "margin": "5% 10%" }} className="ui modal">
                <i className="close icon" onClick={() => this.props.modalRef.current.style.display = "none"}></i>
                <div className="header">
                    Shopping Cart
                </div>
                
                <ModalTable total={this.props.total} products={this.props.products} list={this.props.list}/>

                <div className="actions">
                    <div className="ui black deny button" onClick={() => this.props.modalRef.current.style.display = "none"}>
                        Go Back
                    </div>
                    <div className="ui positive right labeled icon button" onClick={() => {
                        this.props.modalRef.current.style.display = "none"
                        alert('Thank You for shopping with us ...')
                    }}>
                        Buy
                        <i className="checkmark icon"></i>
                    </div>
                </div>
            </div>
        )
    }
}