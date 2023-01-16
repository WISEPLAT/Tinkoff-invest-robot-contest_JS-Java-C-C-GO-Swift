//
//  IRobotHistoryRouter.swift
//
//  Created by Andrey Vasilev on 18/05/2022.
//  Copyright © 2022 Andrey Vasilev. All rights reserved.
//

import Foundation

protocol IRobotHistoryRouter: IBaseRouter {

    func showDetails(deal: Deal)
    func showChart(deals: [Deal])
    func showConfig(robot: Robot)
}
