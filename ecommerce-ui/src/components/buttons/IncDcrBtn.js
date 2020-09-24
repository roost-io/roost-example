import React from 'react'

export default class IncDcrBtn extends React.Component {
    render() {
        return (
            <div className="ui mini action input">
                <button style={{"width": "30%", "justifyContent":"center"}} className="ui blue icon button" onClick={() => this.props.countHandler(false)}>
                    <i className="minus icon"></i>
                </button>
                <input style={{"width": "40%", "fontSize": "150%"}} className="center aligned" type="number" readOnly value={this.props.count} />
                <button style={{"width": "30%", "justifyContent":"center"}} className="ui blue icon button" onClick={() => this.props.countHandler(true)}>
                    <i className="plus icon"></i>
                </button>
            </div>
        )
    }
}