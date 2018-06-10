package main

// whalepost
// Copyright (C) 2018 Maximilian Pachl

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// ---------------------------------------------------------------------------------------
//  imports
// ---------------------------------------------------------------------------------------

import (
	"net/http"

	"github.com/faryon93/handlers"
	"github.com/gorilla/mux"
)

// ---------------------------------------------------------------------------------------
//  public functions
// ---------------------------------------------------------------------------------------

// Routes registers all routes for this endpoint.
func Routes(r *mux.Router) {
	r.Methods(http.MethodPut).Path("/service/{ServiceId}").
		Handler(handlers.ChainFunc(ServiceUpdate, handlers.Keyed(Token)))
}
