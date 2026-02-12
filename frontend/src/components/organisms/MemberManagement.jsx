import { useState } from "react";
import Button from "../atoms/Button";
import Input from "../atoms/Input";
import Badge from "../atoms/Badge";
import Select from "../atoms/Select";

const MemberManagement = ({
  members,
  onInviteMember,
  onUpdateRole,
  onRemoveMember,
  isOwner,
}) => {
  const [email, setEmail] = useState("");
  const [isInviting, setIsInviting] = useState(false);

  const handleInvite = async (e) => {
    e.preventDefault();
    setIsInviting(true);
    await onInviteMember(email);
    setEmail("");
    setIsInviting(false);
  };

  const roleOptions = [
    { value: "member", label: "Member" },
    { value: "admin", label: "Admin" },
  ];

  return (
    <div className="bg-white rounded-lg shadow-md p-6">
      <h3 className="text-xl font-semibold mb-4">Team Members</h3>

      {isOwner && (
        <form onSubmit={handleInvite} className="flex gap-2 mb-6">
          <Input
            type="email"
            name="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="Enter user ID to invite"
            required
          />
          <Button type="submit" variant="primary" disabled={isInviting}>
            {isInviting ? "Inviting..." : "Invite"}
          </Button>
        </form>
      )}

      <div className="space-y-3">
        {members.map((member) => (
          <div
            key={member.user_id}
            className="flex flex-col md:flex-row justify-between items-start md:items-center gap-3 p-4 bg-gray-50 rounded-lg"
          >
            <div>
              <p className="font-medium text-gray-800">{member.fullname}</p>
              <p className="text-sm text-gray-600">{member.email}</p>
            </div>
            <div className="flex items-center gap-3">
              <Badge variant={member.role === "admin" ? "primary" : "default"}>
                {member.role}
              </Badge>
              {isOwner && member.role !== "owner" && (
                <div className="flex gap-2">
                  <select
                    className="px-3 py-1 border rounded text-sm"
                    value={member.role}
                    onChange={(e) =>
                      onUpdateRole(member.user_id, e.target.value)
                    }
                  >
                    <option value="member">Member</option>
                    <option value="admin">Admin</option>
                  </select>
                  <Button
                    variant="danger"
                    onClick={() => onRemoveMember(member.user_id)}
                  >
                    Remove
                  </Button>
                </div>
              )}
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default MemberManagement;
