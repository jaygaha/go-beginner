<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Ticketing System</title>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
</head>
<body>
    <div class="container mt-5">
        <h1 class="mb-4">Ticketing System</h1>

        <!-- Create Ticket Form -->
        <h3>Create Ticket</h3>
        <form id="ticketForm" class="mb-4">
            <div class="mb-3">
                <label for="title" class="form-label">Title</label>
                <input type="text" class="form-control" id="title" required>
            </div>
            <div class="mb-3">
                <label for="description" class="form-label">Description</label>
                <textarea class="form-control" id="description" required></textarea>
            </div>
            <div class="mb-3">
                <label for="status" class="form-label">Status</label>
                <select class="form-select" id="status" required>
                    <option value="open">Open</option>
                    <option value="in-progress">In Progress</option>
                    <option value="closed">Closed</option>
                </select>
            </div>
            <div class="mb-3">
                <label for="createdBy" class="form-label">Created By</label>
                <input type="text" class="form-control" id="createdBy" required>
            </div>
            <button type="submit" class="btn btn-primary">Create Ticket</button>
        </form>

        <!-- Update Ticket Form (Hidden by default) -->
        <h3 id="updateFormTitle" class="d-none">Update Ticket</h3>
        <form id="updateTicketForm" class="mb-4 d-none">
            <input type="hidden" id="updateId">
            <div class="mb-3">
                <label for="updateTitle" class="form-label">Title</label>
                <input type="text" class="form-control" id="updateTitle" required>
            </div>
            <div class="mb-3">
                <label for="updateDescription" class="form-label">Description</label>
                <textarea class="form-control" id="updateDescription" required></textarea>
            </div>
            <div class="mb-3">
                <label for="updateStatus" class="form-label">Status</label>
                <select class="form-select" id="updateStatus" required>
                    <option value="open">Open</option>
                    <option value="in-progress">In Progress</option>
                    <option value="closed">Closed</option>
                </select>
            </div>
            <div class="mb-3">
                <label for="updateCreatedBy" class="form-label">Created By</label>
                <input type="text" class="form-control" id="updateCreatedBy" readonly>
            </div>
            <button type="submit" class="btn btn-success">Update Ticket</button>
            <button type="button" class="btn btn-secondary" id="cancelUpdate">Cancel</button>
        </form>

        <!-- Tickets Table -->
        <h3>Tickets</h3>
        <table class="table table-striped">
            <thead>
                <tr>
                    <th>ID</th>
                    <th>Title</th>
                    <th>Description</th>
                    <th>Status</th>
                    <th>Created By</th>
                    <th>Created At</th>
                    <th>Updated At</th>
                    <th>Actions</th>
                </tr>
            </thead>
            <tbody id="ticketTableBody"></tbody>
        </table>
    </div>

    <script>
        // Fetch and display all tickets
        async function fetchTickets() {
            const response = await fetch('/tickets');
            const tickets = await response.json();
            const tbody = document.getElementById('ticketTableBody');
            tbody.innerHTML = '';
            tickets.forEach(ticket => {
                const row = document.createElement('tr');
                row.innerHTML = `
                    <td>${ticket.id}</td>
                    <td>${ticket.title}</td>
                    <td>${ticket.description}</td>
                    <td>${ticket.status}</td>
                    <td>${ticket.created_by}</td>
                    <td>${new Date(ticket.created_at).toLocaleString()}</td>
                    <td>${new Date(ticket.updated_at).toLocaleString()}</td>
                    <td>
                        <button class="btn btn-sm btn-warning" onclick="showUpdateForm(${ticket.id}, '${ticket.title}', '${ticket.description}', '${ticket.status}', '${ticket.created_by}')">Edit</button>
                        <button class="btn btn-sm btn-danger" onclick="deleteTicket(${ticket.id})">Delete</button>
                    </td>
                `;
                tbody.appendChild(row);
            });
        }

        // Create a new ticket
        document.getElementById('ticketForm').addEventListener('submit', async (e) => {
            e.preventDefault();
            const ticket = {
                title: document.getElementById('title').value,
                description: document.getElementById('description').value,
                status: document.getElementById('status').value,
                created_by: document.getElementById('createdBy').value
            };
            await fetch('/tickets', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(ticket)
            });
            document.getElementById('ticketForm').reset();
            fetchTickets();
        });

        // Show update form with ticket data
        function showUpdateForm(id, title, description, status, createdBy) {
            document.getElementById('updateId').value = id;
            document.getElementById('updateTitle').value = title;
            document.getElementById('updateDescription').value = description;
            document.getElementById('updateStatus').value = status;
            document.getElementById('updateCreatedBy').value = createdBy;
            document.getElementById('updateFormTitle').classList.remove('d-none');
            document.getElementById('updateTicketForm').classList.remove('d-none');
            document.getElementById('ticketForm').classList.add('d-none');
        }

        // Update a ticket
        document.getElementById('updateTicketForm').addEventListener('submit', async (e) => {
            e.preventDefault();
            const id = document.getElementById('updateId').value;
            const ticket = {
                title: document.getElementById('updateTitle').value,
                description: document.getElementById('updateDescription').value,
                status: document.getElementById('updateStatus').value
                // created_by is readonly, not sent in update
            };
            await fetch(`/tickets/${id}`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(ticket)
            });
            document.getElementById('updateTicketForm').reset();
            document.getElementById('updateTicketForm').classList.add('d-none');
            document.getElementById('updateFormTitle').classList.add('d-none');
            document.getElementById('ticketForm').classList.remove('d-none');
            fetchTickets();
        });

        // Cancel update
        document.getElementById('cancelUpdate').addEventListener('click', () => {
            document.getElementById('updateTicketForm').reset();
            document.getElementById('updateTicketForm').classList.add('d-none');
            document.getElementById('updateFormTitle').classList.add('d-none');
            document.getElementById('ticketForm').classList.remove('d-none');
        });

        // Delete a ticket
        async function deleteTicket(id) {
            await fetch(`/tickets/${id}`, { method: 'DELETE' });
            fetchTickets();
        }

        // Initial fetch of tickets
        fetchTickets();
    </script>
</body>
</html>